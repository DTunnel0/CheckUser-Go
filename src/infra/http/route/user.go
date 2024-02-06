package route

import (
	"github.com/DTunnel0/CheckUser-Go/src/infra/adapter"
	"github.com/DTunnel0/CheckUser-Go/src/infra/factory"
	"github.com/labstack/echo/v4"
)

func CreateUserRoute(g *echo.Group) {
	g.GET("/check/:username", adapter.NewEchoAdapter(factory.MakeCheckUserHandler()).Adapt)
	g.GET("/details/:username", adapter.NewEchoAdapter(factory.MakeDetailsUserHandler()).Adapt)
	g.GET("/count", adapter.NewEchoAdapter(factory.MakeCountConnectionsHandler()).Adapt)
}
