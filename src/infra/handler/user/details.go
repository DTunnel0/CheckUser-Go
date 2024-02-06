package user_handler

import (
	"context"

	user_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/user"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
)

type detailUserHandler struct {
	detailUserUseCase *user_use_case.DetailUserUseCase
}

func NewDetailUserHandler(detailUserUseCase *user_use_case.DetailUserUseCase) handler.Handler {
	return &detailUserHandler{detailUserUseCase}
}

func (h *detailUserHandler) Handle(ctx context.Context, request *handler.HttpRequest) (*handler.HttpResponse, error) {
	username := request.Query("username")
	data, err := h.detailUserUseCase.Execute(ctx, username)
	if err != nil {
		return nil, err
	}
	return &handler.HttpResponse{Status: 200, Body: data}, nil
}
