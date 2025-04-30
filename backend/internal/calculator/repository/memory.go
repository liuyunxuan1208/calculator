package repository

import (
	"calculator/backend/internal/calculator/domain"
	"sync"
)

// MemoryRepository 内存存储实现（如需历史记录功能）
type MemoryRepository struct {
	mu      sync.RWMutex
	history []domain.Calculation
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		history: make([]domain.Calculation, 0),
	}
}

// SaveCalculation 保存计算记录
func (r *MemoryRepository) SaveCalculation(op domain.Calculation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.history = append(r.history, op)
	return nil
}

// GetHistory 获取计算历史
func (r *MemoryRepository) GetHistory() ([]domain.Calculation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]domain.Calculation(nil), r.history...), nil
}
