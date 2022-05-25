package pkg

import (
	"go.uber.org/fx"
	"golang-ddd-boilerplate/pkg/fxmysql"
)

var FXPkg = fx.Options(
	fxmysql.FXMysql,
)
