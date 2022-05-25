package requestobject

type GetRoleLimitByRoleIDReq struct {
	RoleID string `json:"roleId" binding:"required"`
}
