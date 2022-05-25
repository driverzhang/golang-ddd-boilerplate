package infrastructure

import (
	"context"
	"git.dustess.com/mp-biz/contrib/errors"
	"go.uber.org/fx"
	"golang-ddd-boilerplate/config"
	"golang-ddd-boilerplate/internal/domain/entity"
	"golang-ddd-boilerplate/internal/domain/repo"
	"golang-ddd-boilerplate/internal/domain/vo"
	"golang-ddd-boilerplate/internal/infrastructure/database"
	"golang-ddd-boilerplate/pkg/fxlog/log"
	"golang-ddd-boilerplate/pkg/fxmysql/gorm"
	gorm2 "gorm.io/gorm"
)

var Provide = fx.Provide(newRoleRepo)

type roleRepo struct {
	log  log.Helper
	conf config.Config
	db   *gorm.DB
}

func newRoleRepo(conf config.Config, log log.Helper, db *gorm.DB) repo.IRoleRepo {
	return &roleRepo{
		log:  log,
		conf: conf,
		db:   db,
	}
}

func (p *roleRepo) withDebug(ctx context.Context) *gorm2.DB {
	return p.db.WithContext(ctx).Debug()
}

func (r roleRepo) get(ctx context.Context, roleID string) (database.Role, error) {
	data := database.Role{}
	err := r.withDebug(ctx).Where("id = ?", roleID).First(&data).Error
	if err != nil {
		if err == gorm2.ErrRecordNotFound {
			return data, errors.New("not found ...")
		}
		return data, err
	}
	return data, nil
}

func (r roleRepo) Get(ctx context.Context, roleID string) (entity.Entity, error) {
	result := entity.Entity{}
	roleData, err := r.get(ctx, roleID)
	if err != nil {
		return result, err
	}

	result = entity.Entity{
		ID:      roleData.ID,
		Account: roleData.Account,
		Name:    roleData.Name,
		Memo:    roleData.Memo,
		Type:    roleData.Type,
		ValueObj: vo.ValueObj{
			PcMenuLimits:   roleData.PcMenuLimits,
			PcFuncLimits:   roleData.PcFuncLimits,
			PcDataLimits:   roleData.PcDataLimits,
			QwMenuLimits:   roleData.QwMenuLimits,
			QwFuncLimits:   roleData.QwFuncLimits,
			QwDataLimits:   roleData.QwDataLimits,
			SCRMMenuLimits: roleData.SCRMMenuLimits,
			SCRMFuncLimits: roleData.SCRMFuncLimits,
			SCRMDataLimits: roleData.SCRMDataLimits,
			ToolLimits:     roleData.ToolLimits,
		},
	}
	return result, nil
}
