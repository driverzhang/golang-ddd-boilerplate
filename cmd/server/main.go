package main

import (
	"go.uber.org/fx"
	"golang-ddd-boilerplate/config"
	"golang-ddd-boilerplate/internal/controller"
	"golang-ddd-boilerplate/internal/domain/service"
	"golang-ddd-boilerplate/internal/infrastructure"
	"golang-ddd-boilerplate/pkg"
	"golang-ddd-boilerplate/pkg/fxgin"
	"golang-ddd-boilerplate/pkg/fxlog"
	"golang-ddd-boilerplate/pkg/fxmysql/gorm"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func init() {
	// 可引入 三方config
	err := config.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	gorm.Init(config.Get().MySQL)
	modules := fx.Options(
		config.Provide,
		fxlog.FXLog,
		service.Provide,
		infrastructure.FXInfra,
		controller.FXController,
		fxgin.FXGin,
		pkg.FXPkg,
	)

	err := fx.ValidateApp(modules)
	if err != nil {
		panic(err)
	}

	app := fx.New(modules)
	app.Run()
}
