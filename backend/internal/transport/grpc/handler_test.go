package grpc_test

import (
	"calculator/backend/internal/calculator/repository"
	"calculator/backend/internal/calculator/service"
	"calculator/backend/internal/transport/grpc"
	calcv1 "calculator/backend/pkg/protobuf/calculator/v1"
	"connectrpc.com/connect"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_Calculate(t *testing.T) {
	// 初始化处理器
	repo := repository.NewMemoryRepository()
	handler := grpc.NewHandler(service.NewCalculatorService(repo))

	t.Run("成功计算", func(t *testing.T) {
		resp, err := handler.Calculate(context.Background(),
			connect.NewRequest(&calcv1.CalculationRequest{
				Operand1: 10,
				Operand2: 2,
				Operator: "*",
			}),
		)
		require.NoError(t, err)
		require.Equal(t, float64(20), resp.Msg.Result)
	})

	t.Run("无效运算符", func(t *testing.T) {
		_, err := handler.Calculate(context.Background(),
			connect.NewRequest(&calcv1.CalculationRequest{
				Operator: "invalid",
			}),
		)
		require.Error(t, err)
		require.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}
