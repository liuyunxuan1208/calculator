package service

import (
	"calculator/backend/internal/calculator/domain"
	"calculator/backend/internal/calculator/repository"
)

type Calculator struct{}

type CalculatorService struct {
	repo *repository.MemoryRepository // 添加存储依赖
}

func NewCalculatorService(repo *repository.MemoryRepository) *CalculatorService {
	return &CalculatorService{repo: repo}
}

func (s *CalculatorService) Calculate(op domain.Calculation) (float64, error) {
	result, err := s.compute(op)
	if err != nil {
		return 0, err
	}

	// 保存计算记录
	_ = s.repo.SaveCalculation(op) // 简单示例忽略错误处理
	return result, nil
}

// GetHistory 获取计算历史记录
func (s *CalculatorService) GetHistory() ([]domain.Calculation, error) {
	return s.repo.GetHistory()
}

func (s *CalculatorService) compute(op domain.Calculation) (float64, error) {
	switch op.Operator {
	case domain.Add:
		return op.Operand1 + op.Operand2, nil
	case domain.Subtract:
		return op.Operand1 - op.Operand2, nil
	case domain.Multiply:
		return op.Operand1 * op.Operand2, nil
	case domain.Divide:
		if op.Operand2 == 0 {
			return 0, domain.ErrDivisionByZero
		}
		return op.Operand1 / op.Operand2, nil
	default:
		return 0, domain.ErrInvalidOperator
	}
}
