package asyncgo_test

import (
	`context`
	`fmt`
	`log`
	"reflect"
	"testing"
	`time`

	"github.com/photowey/asyncgo"
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
			future := asyncgo.Run(tt.args.fx)
			got, err := future.Await()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() -> Await() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() -> Await() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type user struct {
	name string
}

func TestRunz(t *testing.T) {
	type args struct {
		fx      func() any
		factory asyncgo.AwaitFuncFactory
	}
	factory := func(ch chan struct{}, result *any) asyncgo.AwaitFunc {
		return func(ctx context.Context) (any, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-ch:
				return *result, nil
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
				factory: asyncgo.CreateAwaitFunc,
			},
			want:    11111,
			wantErr: false,
		},
		{
			name: "Test future Runz()",
			args: args{
				fx:      findUser,
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
			future := asyncgo.Runz(tt.args.fx, tt.args.factory)
			got, err := future.Await()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Runz() -> Await() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Runz() -> Await() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	factoryCtx := func(ch chan struct{}, result *any) asyncgo.AwaitFunc {
		return func(ctx context.Context) (any, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-ch:
				ctxValue := ctx.Value("asyncgo").(int)
				fmt.Printf("ctxValue = %v", ctxValue)

				return *result, nil
			}
		}
	}

	ctxTests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "Test future Runz()",
			args: args{
				fx:      fn,
				factory: factoryCtx,
			},
			want:    1521,
			wantErr: false,
		},
	}

	for _, tt := range ctxTests {
		t.Run(tt.name, func(t *testing.T) {
			future := asyncgo.Run(tt.args.fx)
			ctx := context.WithValue(context.Background(), "asyncgo", 21)
			got, err := future.Await(ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() -> Await() = %v, want %v", got, tt.want)
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

func fn() any {
	time.Sleep(3 * time.Second)

	return 1521
}

func findUser() any {
	time.Sleep(3 * time.Second)

	return user{
		name: "sharkchili",
	}
}
