package store_api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/nfode/dummy-app/internal/apis/internal/auth"
	"github.com/nfode/dummy-app/internal/apis/store/internal/gen"
)

type server struct {
}

func (server) GetOrderId(_ context.Context, request gen.GetOrderIdRequestObject) (gen.GetOrderIdResponseObject, error) {
	id := request.Id
	item := gen.OrderItem(id)
	price := 1
	return gen.GetOrderId200JSONResponse{
		Id:    &id,
		Item:  &item,
		Price: &price,
	}, nil
}

func (server) PutOrderId(ctx context.Context, request gen.PutOrderIdRequestObject) (gen.PutOrderIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

var _ gen.StrictServerInterface = (*server)(nil)

type StoreAPI struct {
}

func (StoreAPI) RegisterAPI(e *echo.Echo, baseUrl string, middleware auth.Middleware) {
	handler := gen.NewStrictHandler(server{}, []gen.StrictMiddlewareFunc{middleware})
	gen.RegisterHandlersWithBaseURL(e, handler, baseUrl)
}
