package entity

import (
	"golang-ddd-boilerplate/internal/domain/vo"
)

type Entity struct {
	ID      string `json:"id"`               // 主键Id，uuid roleId
	Account string `json:"account"`          // 账户Id
	Name    string `json:"name"`             // 角色名
	Memo    string `json:"memo"`             // 备注
	Type    string `json:"type" bson:"type"` // 角色类型（admin 超级管理员，manage 管理员，sale 销售员）
	vo.ValueObj
}
