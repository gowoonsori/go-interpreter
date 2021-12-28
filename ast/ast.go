package ast

import (
	"bytes"
)

// Node
//let <identifer> = <expresiion>;
type Node interface {
	TokenLiteral() string
	ToString() string
}

//Program AST의 부모 노드
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) ToString() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.ToString())
	}
	return out.String()
}
