package di

import (
	"github.com/google/wire"

	api "github.com/thnkrn/go-fiber-clean-arch/pkg/api"
	middleware "github.com/thnkrn/go-fiber-clean-arch/pkg/api/middleware"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	middleware.NewErrorHandler,
	middleware.NewAuthentication,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)
