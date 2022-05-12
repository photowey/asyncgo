/*
 * Copyright © 2022 photowey (photowey@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package asyncgo_test

import (
	`context`
	`fmt`
	`log`
	`time`

	`github.com/photowey/asyncgo`
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
	future := asyncgo.Run(fx)

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
	factoryCtx := func(ch chan struct{}, result *any) asyncgo.AwaitFunc {
		return func(ctx context.Context) (any, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-ch:
				ctxValue := ctx.Value("asyncgo").(int)
				log.Printf("ctxValue got = %v, want = %v", ctxValue, 21)

				return *result, nil
			}
		}
	}

	// 3.runz -> return future
	future := asyncgo.Runz(fx, factoryCtx)

	// 4.ctx
	ctx := context.WithValue(context.Background(), "asyncgo", 21)
	// 5.await result by sync.
	result, _ := future.Await(ctx)
	fmt.Println(result)

	// Output: 357
}
