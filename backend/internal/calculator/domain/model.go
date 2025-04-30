package domain

import "errors"

type Operation string

const (
	Add      Operation = "+"
	Subtract Operation = "-"
	Multiply Operation = "*"
	Divide   Operation = "/"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrInvalidOperator = errors.New("invalid operator")
)

type Calculation struct {
	Operand1 float64
	Operand2 float64
	Operator Operation
	Result   float64 // 计算结果
}
