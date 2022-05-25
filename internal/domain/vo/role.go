package vo

type ValueObj struct {
	PcMenuLimits   []string `json:"pc_menu_limits"`   // 服务商后台菜单权限
	PcFuncLimits   []string `json:"pc_func_limits"`   // 服务商后台功能权限Id
	PcDataLimits   []string `json:"pc_data_limits"`   // 服务商后台数据权限Id
	QwMenuLimits   []string `json:"qw_menu_limits"`   // 企微应用菜单权限
	QwFuncLimits   []string `json:"qw_func_limits"`   // 企微应用功能权限
	QwDataLimits   []string `json:"qw_data_limits"`   // 企微应用数据权限
	SCRMMenuLimits []string `bson:"scrm_menu_limits"` // scrm助手菜单权限
	SCRMFuncLimits []string `json:"scrm_func_limits"` // scrm助手功能权限
	SCRMDataLimits []string `json:"scrm_data_limits"` // scrm助手数据权限
	ToolLimits     []string `json:"tool_limits"`      // 营销应用中心权限
}
