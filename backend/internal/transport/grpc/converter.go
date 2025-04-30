package grpc

import (
	"calculator/backend/internal/calculator/domain"
	calcv1 "calculator/backend/pkg/protobuf/calculator/v1"
)

// ProtoToDomain 将Protobuf请求转为领域模型
func ProtoToDomain(req *calcv1.CalculationRequest) domain.Calculation {
	return domain.Calculation{
		Operand1: req.Operand1,
		Operand2: req.Operand2,
		Operator: domain.Operation(req.Operator),
	}
}

// DomainToProto 将领域模型转为Protobuf响应
func DomainToProto(result float64) *calcv1.CalculationResponse {
	return &calcv1.CalculationResponse{
		Result: result,
	}
}

// DomainToProtoHistory 将领域模型历史记录转为Protobuf响应
func DomainToProtoHistory(history []domain.Calculation) *calcv1.CalculationHistoryResponse {
	resp := &calcv1.CalculationHistoryResponse{}
	for _, op := range history {
		resp.Items = append(resp.Items, &calcv1.CalculationHistoryItem{
			Operand1: op.Operand1,
			Operand2: op.Operand2,
			Operator: string(op.Operator),
			Result:   op.Result,
		})
	}
	return resp
}

// HistoryToProto 历史记录转换（兼容旧版本）
func HistoryToProto(history []domain.Calculation) *calcv1.CalculationHistoryResponse {
	return DomainToProtoHistory(history)
}
