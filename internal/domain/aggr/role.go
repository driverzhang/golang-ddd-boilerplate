package aggr

import (
	"context"
	"git.dustess.com/mp-biz/contrib/fxpkg/log"
	"golang-ddd-boilerplate/config"
	"golang-ddd-boilerplate/internal/domain/entity"
	"golang-ddd-boilerplate/internal/domain/repo"
)

type AggregateFunc func(aggregate *Aggregate)

type Aggregate struct {
	roleEntity entity.Entity
	roleRope   repo.IRoleRepo
	log        log.Helper
	conf       config.Config
}

func AggWithConfig(conf config.Config) AggregateFunc {
	return func(a *Aggregate) {
		a.conf = conf
	}
}

func AggWithLog(log log.Helper) AggregateFunc {
	return func(a *Aggregate) {
		a.log = log
	}
}

func AggWithRoleRepo(roleRope repo.IRoleRepo) AggregateFunc {
	return func(a *Aggregate) {
		a.roleRope = roleRope
	}
}

// NewRoleAggregateFactory 聚合根工场函数
func NewRoleAggregateFactory(ctx context.Context, roleId string, option ...AggregateFunc) (Aggregate, error) {
	result := Aggregate{}
	for _, v := range option {
		v(&result)
	}

	roleData, err := result.roleRope.Get(ctx, roleId)
	if err != nil {
		return result, err
	}

	result.roleEntity = entity.Entity{
		ID:      roleData.ID,
		Account: roleData.Account,
		Name:    roleData.Name,
		Memo:    roleData.Memo,
	}
	return result, nil
}

func (a *Aggregate) Run() error {
	// do ...
	return nil
}
