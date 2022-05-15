package asyncgox_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/photowey/asyncgo/asyncgox"
)

//
// Example of future
//

func ExampleRun() {
	// 1.fx
	fx := func() any {
		log.Println("async start")
		time.Sleep(3 * time.Second)
		log.Println("async end")

		return 357
	}

	// 2.run -> return future
	future := asyncgox.Run(fx)

	// 3.await result by sync.
	result, _ := future.Await()
	fmt.Println(result)

	// Output: 357
}

func ExampleRunz() {
	// 1.fx
	fx := func() any {
		log.Println("async start")
		time.Sleep(3 * time.Second)
		log.Println("async end")

		return 357
	}

	// 2.factory
	factoryCtx := func(ch chan any) asyncgox.AwaitFunc {
		return func(ctx context.Context) (any, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case result := <-ch:
				defer func() {
					close(ch)
				}()
				ctxValue := ctx.Value("asyncgo").(int)
				log.Printf("ctxValue got = %v, want = %v", ctxValue, 21)

				return result, nil
			}
		}
	}

	// 3.runz -> return future
	future := asyncgox.Runz(fx, factoryCtx)

	// 4.ctx
	ctx := context.WithValue(context.Background(), "asyncgo", 21)
	// 5.await result by sync.
	result, _ := future.Await(ctx)
	fmt.Println(result)

	// Output: 357
}
