package http

import (
	"framework/middleware"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	middleware.ProviderSet,

	NewEcho,
	wire.Struct(new(HTTPServer), "*"),
)
