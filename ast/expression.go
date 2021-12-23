package ast

import "monkey/token"

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
