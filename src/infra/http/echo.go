package http

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/DTunnel0/CheckUser-Go/src/infra/http/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme"
)

func Start(host string, port int, sslEnabled bool) {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, HTML_CONTENT)
	})

	e.GET("/device", func(c echo.Context) error {
		return c.HTML(http.StatusOK, DEVICE_HTML_CONTENT)
	})

	api := e.Group("")
	route.CreateUserRoute(api)
	route.CreateDeviceRoute(api)

	addr := fmt.Sprintf("%s:%d", host, port)

	if !sslEnabled {
		e.Logger.Fatal(e.Start(addr))
		return
	}

	certificate, err := tls.X509KeyPair([]byte(CERT_CONTENT), []byte(KEY_CONTENT))
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	server := http.Server{
		Addr:    addr,
		Handler: e,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{certificate},
			NextProtos:   []string{acme.ALPNProto},
		},
	}

	if err := server.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
