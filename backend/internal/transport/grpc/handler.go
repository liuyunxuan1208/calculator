package grpc

import (
	"calculator/backend/internal/calculator/service"
	calcv1 "calculator/backend/pkg/protobuf/calculator/v1"
	"github.com/bufbuild/connect-go"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc *service.CalculatorService
}

func NewHandler(svc *service.CalculatorService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Calculate(
	ctx context.Context,
	req *connect.Request[calcv1.CalculationRequest],
) (*connect.Response[calcv1.CalculationResponse], error) {
	// 使用转换器
	op := ProtoToDomain(req.Msg)

	result, err := h.svc.Calculate(op)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(DomainToProto(result)), nil
}

func (h *Handler) GetHistory(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[calcv1.CalculationHistoryResponse], error) {
	history, err := h.svc.GetHistory()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(DomainToProtoHistory(history)), nil
}
