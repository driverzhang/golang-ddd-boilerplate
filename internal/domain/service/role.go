/**
 * @Description: 若你不想写DDD，也可以直接在service里写你的业务逻辑
 */
package service

import (
	"context"
	"git.dustess.com/mp-biz/contrib/fxpkg/log"
	"go.uber.org/fx"
	"golang-ddd-boilerplate/config"
	"golang-ddd-boilerplate/internal/domain/aggr"
	"golang-ddd-boilerplate/internal/domain/repo"
)

var Provide = fx.Provide(newRoleService)

type Service struct {
	roleRope repo.IRoleRepo
	log      log.Helper
	conf     config.Config
}

func newRoleService(conf config.Config, log log.Helper, roleRope repo.IRoleRepo) Service {
	return Service{
		roleRope: roleRope,
		log:      log,
		conf:     conf,
	}
}

func (s *Service) GetSome(ctx context.Context, id string) error {
	roleAgg, err := aggr.NewRoleAggregateFactory(ctx, id,
		aggr.AggWithConfig(s.conf),
		aggr.AggWithLog(s.log),
		aggr.AggWithRoleRepo(s.roleRope))
	if err != nil {
		return err
	}
	return roleAgg.Run()
}
