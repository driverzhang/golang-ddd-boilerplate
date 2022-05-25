package viewobject

import (
	"golang-ddd-boilerplate/internal/domain/entity"
)

type Role struct {
	ID      string `json:"id" bson:"_id"`          // 主键Id，uuid
	Account string `json:"account" bson:"account"` // 账户Id
	Name    string `json:"name" bson:"name"`       // 角色名
	Type    string `json:"type" bson:"type"`       // 角色类型（admin 超级管理员，manage 管理员，sale 销售员）
}

type GetRoleResp struct {
	Role
}

func NewGetRoleResp(role entity.Entity) GetRoleResp {
	return GetRoleResp{Role{
		ID:      role.ID,
		Account: role.Account,
		Name:    role.Name,
		Type:    role.Type,
	}}
}
