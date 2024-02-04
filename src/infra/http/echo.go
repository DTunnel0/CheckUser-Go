package http

import (
	"fmt"
	"net/http"

	"github.com/DTunnel0/CheckUser-Go/src/infra/http/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(host string, port int) {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, HTML_CONTENT)
	})

	api := e.Group("")
	route.CreateUserRoute(api)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
