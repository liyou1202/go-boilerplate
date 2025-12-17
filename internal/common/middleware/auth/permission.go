package auth

import (
	"sync_drive_backend/internal/common/consts"
	"sync_drive_backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

// RequireRole 權限控制中介層：要求特定角色
func RequireRole(allowedRoles ...consts.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := GetRoleID(c)
		if roleID == "" {
			errors.HandleError(c, errors.New(errors.ErrUnauthorized, "user role not found"))
			c.Abort()
			return
		}

		// 檢查角色是否被允許
		userRole := consts.Role(roleID)
		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		// 無權限
		errors.HandleError(c, errors.New(errors.ErrUnauthorized, "insufficient permissions"))
		c.Abort()
	}
}

// RequireAdmin 權限控制中介層：要求管理員權限
func RequireAdmin() gin.HandlerFunc {
	return RequireRole(consts.RoleAdmin)
}
