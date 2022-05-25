package database

type Role struct {
	ID             string   `json:"id" bson:"_id"`                            // 主键Id，uuid
	Account        string   `json:"account" bson:"account"`                   // 账户Id
	Name           string   `json:"name" bson:"name"`                         // 角色名
	Memo           string   `json:"memo" bson:"memo"`                         // 备注
	OwnNum         int      `json:"own_num" bson:"own_num"`                   // 成员数量
	PcMenuLimits   []string `json:"pc_menu_limits" bson:"pc_menu_limits"`     // 服务商后台菜单权限
	PcFuncLimits   []string `json:"pc_func_limits" bson:"pc_func_limits"`     // 服务商后台功能权限Id
	PcDataLimits   []string `json:"pc_data_limits" bson:"pc_data_limits"`     // 服务商后台数据权限Id
	QwMenuLimits   []string `json:"qw_menu_limits" bson:"qw_menu_limits"`     // 企微应用菜单权限
	QwFuncLimits   []string `json:"qw_func_limits" bson:"qw_func_limits"`     // 企微应用功能权限
	QwDataLimits   []string `json:"qw_data_limits" bson:"qw_data_limits"`     // 企微应用数据权限
	SCRMMenuLimits []string `bson:"scrm_menu_limits"`                         // scrm助手菜单权限
	SCRMFuncLimits []string `json:"scrm_func_limits" bson:"scrm_func_limits"` // scrm助手功能权限
	SCRMDataLimits []string `json:"scrm_data_limits" bson:"scrm_data_limits"` // scrm助手数据权限
	Type           string   `json:"type" bson:"type"`                         // 角色类型（admin 超级管理员，manage 管理员，sale 销售员）
	Status         string   `json:"status" bson:"status"`                     // 状态（use：在用，stop：停用）
	Creator        string   `json:"creator" bson:"creator"`                   // 创建者
	CreateTime     string   `json:"create_time" bson:"create_time"`           // 创建时间
	LastTime       string   `json:"last_time" bson:"last_time"`               // 最新修改时间
	Level          int8     `json:"level" bson:"level"`                       // 角色等级
	Resources      []string `json:"resourceIds" bson:"resourceIds"`           // 资源数组
	ChildID        []string `json:"childIds" bson:"childIds"`                 // 下一级别的角色id数组
	ToolLimits     []string `json:"tool_limits" bson:"tool_limits"`           // 营销应用中心权限
}
