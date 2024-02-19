package service

import (
	"context"
	"strings"
)

type (
	Service struct{}

	Request struct {
		Message  string `json:"message"`
		Username string `json:"username"`
	}

	Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)

func New() *Service {
	return &Service{}
}

func (a *Service) Service(ctx context.Context, in *Request) (*Response, error) {

	if strings.Contains(in.Message, "error") {
		return &Response{
			Status:  "500",
			Message: "error",
		}, nil
	}

	return &Response{
		Status:  "OK",
		Message: "todo ok",
	}, nil
}
