// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ratelimit

import "context"

type Limiter interface {
	// Limit 用于判断是否触发限流. key 表示限流的键值，比如用户 ID 等.
	// bool 表示是否限流，true 表示限流了
	// error 表示判断是否限流时发生了错误
	Limit(ctx context.Context, key string) (bool, error)
}
