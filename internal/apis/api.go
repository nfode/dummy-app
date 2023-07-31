package apis

import (
	"github.com/labstack/echo/v4"
	"github.com/nfode/dummy-app/internal/apis/store"
)

type RegisterAPI interface {
	RegisterAPI(e *echo.Echo, baseUrl string)
}

type API struct {
	Echo *echo.Echo
}

var baseUrl = "/api"

func (a API) registerAPIs(apis ...RegisterAPI) {
	for _, api := range apis {
		api.RegisterAPI(a.Echo, baseUrl)
	}
}

func (a API) Setup() {
	a.registerAPIs(
		store_api.StoreAPI{},
	)
}
