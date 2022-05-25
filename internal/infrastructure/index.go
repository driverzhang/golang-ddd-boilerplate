package infrastructure

import (
	"go.uber.org/fx"
)

var FXInfra = fx.Options(
	Provide,
)
