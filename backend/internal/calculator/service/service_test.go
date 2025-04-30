package service_test

import (
	"calculator/backend/internal/calculator/domain"
	"calculator/backend/internal/calculator/repository"
	"calculator/backend/internal/calculator/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatorService_Calculate(t *testing.T) {
	// 初始化依赖（使用内存仓库）
	repo := repository.NewMemoryRepository()
	svc := service.NewCalculatorService(repo)

	tests := []struct {
		name     string
		op       domain.Calculation
		expected float64
		wantErr  bool
	}{
		{"加法", domain.Calculation{Operand1: 2, Operand2: 3, Operator: domain.Add, Result: 0}, 5, false},
		{"减法", domain.Calculation{Operand1: 5, Operand2: 3, Operator: domain.Subtract, Result: 0}, 2, false},
		{"乘法", domain.Calculation{Operand1: 2, Operand2: 3, Operator: domain.Multiply, Result: 0}, 6, false},
		{"除法", domain.Calculation{Operand1: 6, Operand2: 2, Operator: domain.Divide, Result: 0}, 3, false},
		{"除法零错误", domain.Calculation{Operand1: 1, Operand2: 0, Operator: domain.Divide, Result: 0}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.Calculate(tt.op)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}
