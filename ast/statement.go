package ast

import "monkey/token"

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token //let
	Name  *Identifier //변수 바인딩 식별자
	Value Expression  //값을 생성하는 표현식
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
