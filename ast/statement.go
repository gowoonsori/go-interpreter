package ast

import (
	"bytes"
	"monkey/token"
)

type Statement interface {
	Node
	statementNode()
}

//LetStatement let 구문
type LetStatement struct {
	Token token.Token //let
	Name  *Identifier //변수 바인딩 식별자
	Value Expression  //값을 생성하는 표현식
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) toString() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.toString())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.toString())
	}
	out.WriteString(";")

	return out.String()
}

//ReturnStatement return 구문
type ReturnStatement struct {
	Token       token.Token //return
	ReturnValue Expression  //반환값은 값일 수도 표현식일 수도 있다.
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) toString() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.toString())
	}
	out.WriteString(";")

	return out.String()
}

//ExpressionStatement 표현 구문
//let과 return구문을 제외한 모든 구문
type ExpressionStatement struct {
	Token      token.Token //표현식의 첫번째 토큰
	Expression Expression  //표현식
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) toString() string {
	if es.Expression != nil {
		return es.Expression.toString()
	}
	return ""
}
