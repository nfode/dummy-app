package store_api

import (
	"context"
	api "github.com/nfode/dummy-app/internal/apis/store/internal/gen"
)

type server struct {
}

func (server) GetOrderId(_ context.Context, request api.GetOrderIdRequestObject) (api.GetOrderIdResponseObject, error) {
	id := request.Id
	item := api.OrderItem(id)
	price := 1
	return api.GetOrderId200JSONResponse{
		Id:    &id,
		Item:  &item,
		Price: &price,
	}, nil
}

func (server) PutOrderId(ctx context.Context, request api.PutOrderIdRequestObject) (api.PutOrderIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

var _ api.StrictServerInterface = (*server)(nil)

func RegisterStoreAPI(e api.EchoRouter) {
	handler := api.NewStrictHandler(server{}, nil)
	api.RegisterHandlers(e, handler)
}
