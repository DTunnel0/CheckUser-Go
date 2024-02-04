package adapter

import (
	"net/http"

	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
	"github.com/labstack/echo/v4"
)

const (
	badRequest    = http.StatusBadRequest
	internalError = http.StatusInternalServerError
)

type EchoAdapter struct {
	handler handler.Handler
}

func NewEchoAdapter(handler handler.Handler) *EchoAdapter {
	return &EchoAdapter{handler: handler}
}

func newResponse(status int, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": status,
		"data":   data,
	}
}

func newErrorResponse(status int, message string) map[string]interface{} {
	return newResponse(status, map[string]string{"error": message})
}

func (ed *EchoAdapter) Adapt(e echo.Context) error {
	query := map[string]interface{}{}
	body := map[string]interface{}{}

	for key, values := range e.QueryParams() {
		if len(values) > 0 {
			query[key] = values[0]
		}
	}

	for _, name := range e.ParamNames() {
		query[name] = e.Param(name)
	}

	if err := e.Bind(&body); err != nil {
		return e.JSON(badRequest, newErrorResponse(badRequest, "Erro ao processar dados do corpo da requisição"))
	}

	httpRequest := handler.NewHttpRequest(query, body)
	response, err := ed.handler.Handle(e.Request().Context(), httpRequest)
	if err != nil {
		return e.JSONPretty(internalError, newErrorResponse(internalError, err.Error()), " ")
	}

	return e.JSONPretty(response.Status, response.Body, " ")
}
