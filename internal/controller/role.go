package controller

import (
	"fmt"
	"git.dustess.com/mk-base/gin-ext/extend"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"golang-ddd-boilerplate/config"
	"golang-ddd-boilerplate/internal/client/requestobject"
	"golang-ddd-boilerplate/internal/client/viewobject"
	"golang-ddd-boilerplate/internal/domain/repo"
	"golang-ddd-boilerplate/pkg/fxlog/log"
)

var RunRoleRouter = fx.Invoke(NewRoleHandler)

type RoleHandler struct {
	roleRepo repo.IRoleRepo
	conf     config.Config
	log      log.Helper
}

func NewRoleHandler(router *gin.Engine, roleRepo repo.IRoleRepo,
	conf config.Config, log log.Helper) {
	roleHandler := &RoleHandler{
		roleRepo: roleRepo,
		conf:     conf,
		log:      log,
	}
	grInternal := fmt.Sprintf("/%s/internal/limit", conf.Domain)
	rCode := router.Group(grInternal)
	rCode.POST("/role/batch/v1", roleHandler.GetRole) // 获取企业角色详情
}

// GetRoleLimitByRoleID 获取企业角色详情
func (r *RoleHandler) GetRole(ctx *gin.Context) {
	rqCtx := ctx.Request.Context()
	params := &requestobject.GetRoleLimitByRoleIDReq{}
	if err := ctx.Bind(params); err != nil {
		r.log.WithContext(ctx).Error(err)
		extend.SendParamParseError(ctx)
		return
	}

	// 获取角色详情包含该角色的
	result, err := r.roleRepo.Get(rqCtx, params.RoleID)
	if err != nil {
		r.log.WithContext(rqCtx).Error(err)
	}

	// 强制要求，必须进行数据转换
	resp := viewobject.NewGetRoleResp(result)
	extend.SendData(ctx, resp, err)
}
