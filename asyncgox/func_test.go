package asyncgox_test

import (
	"context"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/photowey/asyncgo/asyncgox"
)

func TestRun(t *testing.T) {
	type args struct {
		fx func() any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test future Run()",
			args: args{
				fx: runInAsync,
			},
			want:    11111,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := asyncgox.Run(tt.args.fx)
			got, err := f.Await()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() -> Await() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func runInAsync() any {
	log.Println("runInAsync start")
	time.Sleep(3 * time.Second)
	log.Println("runInAsync end")

	return 11111
}

func TestRunz(t *testing.T) {
	type args struct {
		fx      func() any
		factory asyncgox.AwaitFuncFactory
	}

	factory := func(ch chan any) asyncgox.AwaitFunc {
		return func(ctx context.Context) (any, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case result := <-ch:
				defer func() {
					close(ch)
					// ch <- Task{} // panic: send on closed channel
				}()
				return result, nil
			}
		}
	}

	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test future Runz()",
			args: args{
				fx:      runInAsync,
				factory: asyncgox.AwaitFuncFactoryFunc,
			},
			want:    11111,
			wantErr: false,
		},
		{
			name: "Test future Runz()",
			args: args{
				fx:      runInAsync,
				factory: factory,
			},
			want: user{
				name: "sharkchili",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asyncgox.Runz(tt.args.fx, tt.args.factory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Runz() = %v, want %v", got, tt.want)
			}
		})
	}
}

type user struct {
	name string
}

func findUser() any {
	time.Sleep(3 * time.Second)

	return user{
		name: "sharkchili",
	}
}
