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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chenmingyong0423/gkit/syncx"
)

func TestCountLimiter_Limit(t *testing.T) {
	l := CountLimiter{
		Rate:     1,
		Interval: 1000 * time.Millisecond,
		Locker:   syncx.NewMapKeyLock(),
		Limits:   map[string]*RateLimitInfo{},
	}
	testCases := []struct {
		name string

		ctx      context.Context
		key      string
		interval time.Duration

		wantB   bool
		wantErr error
	}{
		{
			name:  "通过",
			ctx:   context.Background(),
			key:   "chenmingyong",
			wantB: false,
		},
		{
			name:  "另一个 key 通过",
			ctx:   context.Background(),
			key:   "chenmingyong0423",
			wantB: false,
		},
		{
			name:     "同一个 key 限流",
			ctx:      context.Background(),
			key:      "chenmingyong",
			interval: 500 * time.Millisecond,
			wantB:    true,
		},
		{
			name:     "下一个窗口通过",
			ctx:      context.Background(),
			key:      "chenmingyong",
			interval: 1100 * time.Millisecond,
			wantB:    false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			<-time.After(tc.interval)
			gotB, err := l.Limit(tc.ctx, tc.key)
			assert.Equal(t, tc.wantB, gotB)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
