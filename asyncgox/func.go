package asyncgox

import (
	"context"
)

// AwaitFunc a func of await
type AwaitFunc func(ctx context.Context) (any, error)

// Run executes the async function
func Run(fx func() any) Future {
	return Runz(fx, CreateAwaitFunc)
}

// Runz executes the async function with custom AwaitFunc factory
func Runz(fx func() any, factory AwaitFuncFactory) Future {
	ch := make(chan any)
	go func() {
		result := fx()
		ch <- result
	}()
	return future{
		await: factory(ch),
	}
}
