package controller

import "go.uber.org/fx"

var FXController = fx.Options(RunRoleRouter, RunRouter)
