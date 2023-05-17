package app

import (
	"github.com/ekkinox/fx-template/app/handlers"
	"github.com/ekkinox/fx-template/modules/fxhttpserver"
	"go.uber.org/fx"
)

func RegisterHandlers() fx.Option {
	return fx.Options(
		fxhttpserver.RegisterHandler("GET", "/foo", handlers.NewFooHandler),
		fxhttpserver.RegisterHandler("GET", "/bar", handlers.NewBarHandler),
	)
}
