package asyncgox

import (
	"context"
)

// AwaitFuncFactory a factory of AwaitFunc
type AwaitFuncFactory func(ch chan any) AwaitFunc

// AwaitFuncFactoryFunc a func of AwaitFuncFactory
func AwaitFuncFactoryFunc(ch chan any) AwaitFunc {
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
