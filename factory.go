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
	"context"
)

// AwaitFuncFactory a factory of AwaitFunc
type AwaitFuncFactory func(ch chan struct{}, result *any) AwaitFunc

// CreateAwaitFunc a func of AwaitFuncFactory
func CreateAwaitFunc(ch chan struct{}, result *any) AwaitFunc {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ch:
			return *result, nil
		}
	}
}
