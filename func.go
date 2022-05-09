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

package asyncgo

import (
	`context`
)

// AwaitFunc {@code Future} {@code Await} func
type AwaitFunc func(ctx context.Context) (any, error)

// Run executes the async function
func Run(fx func() any) Future {
	return Runz(fx, CreateAwaitFunc)
}

// Runz executes the async function with custom {@code AwaitFunc} factory
func Runz(fx func() any, factory AwaitFuncFactory) Future {
	var result any
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		result = fx()
	}() // TODO Goroutine pool?
	return future{
		await: factory(ch, &result),
	}
}
