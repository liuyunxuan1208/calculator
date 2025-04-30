package grpc_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"calculator/backend/internal/calculator/domain"
	"calculator/backend/internal/calculator/repository"
	"calculator/backend/internal/calculator/service"
	"calculator/backend/internal/transport/grpc"
	calcv1 "calculator/backend/pkg/protobuf/calculator/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TestCalculate 测试计算功能
func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		request  *calcv1.CalculationRequest
		expected float64
		hasError bool
	}{
		{"加法测试", &calcv1.CalculationRequest{Operand1: 1, Operand2: 2, Operator: "+"}, 3, false},
		{"减法测试", &calcv1.CalculationRequest{Operand1: 5, Operand2: 3, Operator: "-"}, 2, false},
		{"乘法测试", &calcv1.CalculationRequest{Operand1: 2, Operand2: 3, Operator: "*"}, 6, false},
		{"除法测试", &calcv1.CalculationRequest{Operand1: 6, Operand2: 2, Operator: "/"}, 3, false},
		{"除零错误", &calcv1.CalculationRequest{Operand1: 1, Operand2: 0, Operator: "/"}, 0, true},
		{"无效运算符", &calcv1.CalculationRequest{Operand1: 1, Operand2: 1, Operator: "x"}, 0, true},
	}

	repo := repository.NewMemoryRepository()
	svc := service.NewCalculatorService(repo)
	handler := grpc.NewHandler(svc)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := connect.NewRequest(tt.request)
			resp, err := handler.Calculate(context.Background(), req)

			if tt.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, resp.Msg.Result)
			}
		})
	}
}

// TestGetHistory 测试获取历史记录功能
func TestGetHistory(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewCalculatorService(repo)
	handler := grpc.NewHandler(svc)

	// 添加测试数据
	_ = repo.SaveCalculation(domain.Calculation{Operand1: 1, Operand2: 2, Operator: "+"})
	_ = repo.SaveCalculation(domain.Calculation{Operand1: 3, Operand2: 4, Operator: "*"})

	t.Run("获取历史记录", func(t *testing.T) {
		req := connect.NewRequest(&emptypb.Empty{})
		resp, err := handler.GetHistory(context.Background(), req)

		require.NoError(t, err)
		assert.Equal(t, 2, len(resp.Msg.Items))
	})
}