package apis

import (
	"github.com/labstack/echo/v4"
	"github.com/nfode/dummy-app/internal/apis/internal/auth"
	"github.com/nfode/dummy-app/internal/apis/store"
)

type RegisterAPI interface {
	RegisterAPI(e *echo.Echo, baseUrl string, authMiddleware auth.Middleware)
}

type API struct {
	Echo *echo.Echo
}

var baseUrl = "/api"

func (a API) registerAPIs(apis ...RegisterAPI) {
	middleware := auth.MiddlewareFunc()

	for _, api := range apis {
		api.RegisterAPI(a.Echo, baseUrl, middleware)

	}
}

func (a API) Setup() {
	a.registerAPIs(
		store_api.StoreAPI{},
	)
}
