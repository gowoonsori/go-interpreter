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

func (i *Identifier) ToString() string {
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
func (il *IntegerLiteral) ToString() string {
	return il.Token.Literal
}

//PrefixExpression
type PrefixExpression struct {
	Token    token.Token //전위 연산자 토큰 ( ! - + )
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) ToString() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.ToString())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

//InfixExpression
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) ToString() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.ToString())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.ToString())
	out.WriteString(")")

	return out.String()
}

func (ie InfixExpression) expressionNode() {}

//Boolean
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

func (b *Boolean) ToString() string { return b.Token.Literal }

func (b *Boolean) expressionNode() {}
