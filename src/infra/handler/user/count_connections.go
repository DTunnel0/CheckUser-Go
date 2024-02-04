package user_handler

import (
	"context"

	user_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/user"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
)

type countConnectionsHandler struct {
	countConnectionsUseCase *user_use_case.CountConnectionsUseCase
}

func NewCountConnectionsHandler(countConnectionsUseCase *user_use_case.CountConnectionsUseCase) handler.Handler {
	return &countConnectionsHandler{
		countConnectionsUseCase: countConnectionsUseCase,
	}
}

func (h *countConnectionsHandler) Handle(ctx context.Context, request *handler.HttpRequest) (*handler.HttpResponse, error) {
	count := h.countConnectionsUseCase.Execute(ctx)
	return &handler.HttpResponse{
		Body:   count,
		Status: 200,
	}, nil
}
