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
	"context"
	"time"

	"github.com/chenmingyong0423/gkit/syncx"
)

type RateLimitInfo struct {
	// 当前窗口的请求数
	count int
	// 最后重置的时间
	lastReset time.Time
}

type CountLimiter struct {
	// 阈值
	Rate int
	// 窗口大小
	Interval time.Duration
	// 处理键的特定锁
	Locker syncx.KeyLocker
	// 限流信息
	// fixme 后续使用支持对 key 添加过期时间的 local cache 来替代 map
	Limits map[string]*RateLimitInfo
}

func (l *CountLimiter) Limit(_ context.Context, key string) (bool, error) {
	l.Locker.Lock(key)
	defer l.Locker.Unlock(key)
	now := time.Now()
	info, ok := l.Limits[key]
	if !ok {
		l.Limits[key] = &RateLimitInfo{count: 1, lastReset: now}
		return false, nil
	}
	if info.lastReset.Add(l.Interval).Before(now) {
		info.count = 1
		info.lastReset = now
		return false, nil
	} else if info.count < l.Rate {
		info.count++
		return false, nil
	}
	return true, nil
}
