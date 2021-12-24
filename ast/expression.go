package ast

import (
	"bytes"
	"monkey/token"
)

type Expression interface {
	Node
	expressionNode()
}

//Identifier
type Identifier struct {
	Token token.Token //IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) toString() string {
	return i.Value
}

//IntegerLiteral
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) toString() string {
	return il.Token.Literal
}

//PrefixExpression
type PrefixExpression struct {
	Token    token.Token //전위 연산자 토큰 ( ! - + )
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) toString() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.toString())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}
