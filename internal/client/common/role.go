/**
 * @Description: 包名无法写成 关键字const
 */
package common

// RoleType type
type RoleType string

// RoleType
const (
	Admin  RoleType = "admin"  // 超级管理员
	Manage RoleType = "manage" // 管理员
	Sale   RoleType = "sale"   // 销售员
)

const (
	PresetType = "系统默认"
)

const (
	UserRoleStatusUse  = "use"
	UserRoleStatusStop = "stop"
)

const (
	UserRoleTypeNormal = "normal"
)

const (
	UserRoleTypeAdmin  = "admin"
	UserRoleTypeManage = "manage"
)

const (
	MenuLimit = "menu" // 权限类型
	FuncLimit = "func"
	DataLimit = "data"
	ToolLimit = "tool"
)
