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

package local

import (
	"time"

	"github.com/chenmingyong0423/gkit/syncx"

	"github.com/chenmingyong0423/go-ratelimiter/internal/ratelimit/local"
)

// NewCounterLimiter 返回一个 local.CountLimiter 的新实例
// rate 表示限流的阈值
// interval 表示限流的时间窗口
// locker 表示处理键的特定锁
func NewCounterLimiter(rate int, interval time.Duration, locker syncx.KeyLocker) *local.CountLimiter {
	return &local.CountLimiter{
		Rate:     rate,
		Interval: interval,
		Locker:   locker,
		Limits:   map[string]*local.RateLimitInfo{},
	}
}

// NewCounterLimiterWithDefaultLocker 返回一个 local.CountLimiter 的新实例
// rate 表示限流的阈值
// interval 表示限流的时间窗口
func NewCounterLimiterWithDefaultLocker(rate int, interval time.Duration) *local.CountLimiter {
	return &local.CountLimiter{
		Rate:     rate,
		Interval: interval,
		Locker:   syncx.NewMapKeyLock(),
		Limits:   map[string]*local.RateLimitInfo{},
	}
}
