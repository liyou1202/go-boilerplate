package auth

import (
	"strings"

	"sync_drive_backend/pkg/errors"
	"sync_drive_backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	// ContextKeyUserID 用戶 ID 的 context key
	ContextKeyUserID = "user_id"
	// ContextKeyUsername 用戶名稱的 context key
	ContextKeyUsername = "username"
	// ContextKeyRoleID 角色 ID 的 context key
	ContextKeyRoleID = "role_id"
)

// JWTAuth JWT 驗證中介層
func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 Header 取得 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.HandleError(c, errors.New(errors.ErrUnauthorized, "missing authorization header"))
			c.Abort()
			return
		}

		// 驗證格式：Bearer {token}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errors.HandleError(c, errors.New(errors.ErrUnauthorized, "invalid authorization format"))
			c.Abort()
			return
		}

		token := parts[1]

		// 解析 Token
		claims, err := jwt.Parse(token, secret)
		if err != nil {
			errors.HandleError(c, errors.New(errors.ErrUnauthorized, "invalid or expired token"))
			c.Abort()
			return
		}

		// 將用戶資訊存入 Context
		c.Set(ContextKeyUserID, claims.UserID)
		c.Set(ContextKeyUsername, claims.Username)
		c.Set(ContextKeyRoleID, claims.RoleId)

		c.Next()
	}
}

// GetUserID 從 Context 取得用戶 ID
func GetUserID(c *gin.Context) string {
	if userID, exists := c.Get(ContextKeyUserID); exists {
		return userID.(string)
	}
	return ""
}

// GetUsername 從 Context 取得用戶名稱
func GetUsername(c *gin.Context) string {
	if username, exists := c.Get(ContextKeyUsername); exists {
		return username.(string)
	}
	return ""
}

// GetRoleID 從 Context 取得角色 ID
func GetRoleID(c *gin.Context) string {
	if roleID, exists := c.Get(ContextKeyRoleID); exists {
		return roleID.(string)
	}
	return ""
}
