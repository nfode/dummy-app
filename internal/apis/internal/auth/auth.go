package auth

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

type Middleware = runtime.StrictEchoMiddlewareFunc

func MiddlewareFunc() Middleware {
	return func(f runtime.StrictEchoHandlerFunc, operationID string) runtime.StrictEchoHandlerFunc {
		return func(ctx echo.Context, request interface{}) (response interface{}, err error) {
			_ = Checker(ctx)
			return f(ctx, request)
		}
	}
}

func Checker(ctx echo.Context) error {
	// TODO work with proper JWT token and check scopes
	scopes, ok := ctx.Get(scopes).([]string)
	if !ok {
		panic("error something went wrong")
	}
	fmt.Printf("%v", scopes)
	return nil
}

const scopes = "bearerAuth.Scopes"
