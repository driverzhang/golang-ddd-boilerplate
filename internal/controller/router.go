package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var RunRouter = fx.Invoke(Router)

func Router(router *gin.Engine) {
	// 自行建立 独立or功能router
}
