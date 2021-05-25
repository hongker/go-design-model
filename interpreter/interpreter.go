package interpreter

import "strings"

type Expression interface {
	Interpreter(context string) bool
}

type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{data: data}
}

func (expression TerminalExpression) Interpreter(context string) bool {
	if strings.Contains(context, expression.data) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (expression *OrExpression) Interpreter(context string) bool {
	return expression.expr1.Interpreter(context) || expression.expr2.Interpreter(context)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		expr1: expr1,
		expr2: expr2,
	}
}

func (expression *AndExpression) Interpreter(context string) bool {
	return expression.expr1.Interpreter(context) && expression.expr2.Interpreter(context)
}

