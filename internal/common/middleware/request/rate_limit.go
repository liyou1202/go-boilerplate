package request

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 簡單的記憶體限流器
type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int           // 時間窗口內的最大請求數
	window   time.Duration // 時間窗口大小
}

// NewRateLimiter 創建限流器
// limit: 時間窗口內允許的最大請求數
// window: 時間窗口大小（例如：1 分鐘）
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// 定期清理過期的記錄
	go rl.cleanup()

	return rl
}

// RateLimit 限流中介層
func (rl *RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用客戶端 IP 作為限流 key
		key := c.ClientIP()

		if !rl.allow(key) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// allow 檢查是否允許請求
func (rl *RateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window)

	// 取得該 key 的請求記錄
	requests := rl.requests[key]

	// 過濾掉時間窗口外的請求
	var validRequests []time.Time
	for _, req := range requests {
		if req.After(windowStart) {
			validRequests = append(validRequests, req)
		}
	}

	// 檢查是否超過限制
	if len(validRequests) >= rl.limit {
		return false
	}

	// 記錄本次請求
	validRequests = append(validRequests, now)
	rl.requests[key] = validRequests

	return true
}

// cleanup 定期清理過期的記錄
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.window)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		windowStart := now.Add(-rl.window)

		for key, requests := range rl.requests {
			var validRequests []time.Time
			for _, req := range requests {
				if req.After(windowStart) {
					validRequests = append(validRequests, req)
				}
			}

			if len(validRequests) == 0 {
				delete(rl.requests, key)
			} else {
				rl.requests[key] = validRequests
			}
		}
		rl.mu.Unlock()
	}
}
