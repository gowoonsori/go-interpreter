package ast

import "monkey/token"

type Expression interface {
	Node
	expressionNode()
}

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
