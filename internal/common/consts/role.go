package consts

// Role 角色類型
type Role string

const (
	RoleAdmin   Role = "admin"   // 系統管理員
	RoleUser    Role = "user"    // 一般用戶
	RoleVehicle Role = "vehicle" // 車輛
	RoleDevice  Role = "device"  // 設備
)

// IsValid 驗證角色是否有效
func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleUser, RoleVehicle, RoleDevice:
		return true
	default:
		return false
	}
}

// String 返回角色字串
func (r Role) String() string {
	return string(r)
}
