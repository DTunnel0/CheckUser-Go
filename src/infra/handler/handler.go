package handler

import (
	"context"
)

type Handler interface {
	Handle(ctx context.Context, request *HttpRequest) (*HttpResponse, error)
}

type HttpRequest struct {
	query map[string]interface{}
	body  map[string]interface{}
}

func NewHttpRequest(query, body map[string]interface{}) *HttpRequest {
	return &HttpRequest{query, body}
}

func (r *HttpRequest) Body(fieldName string) string {
	val, ok := r.body[fieldName]
	if !ok {
		return ""
	}
	return val.(string)
}

func (r *HttpRequest) Query(paramName string) string {
	val, ok := r.query[paramName]
	if !ok {
		return ""
	}
	return val.(string)
}

type HttpResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}
