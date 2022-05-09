package asyncgo_test

import (
	"reflect"
	"testing"

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
				fx: RunAsync,
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

func TestRunz(t *testing.T) {
	type args struct {
		fx      func() any
		factory asyncgo.AwaitFuncFactory
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
				fx:      RunAsync,
				factory: asyncgo.CreateAwaitFunc,
			},
			want:    11111,
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
}
