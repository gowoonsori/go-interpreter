package ast

// Node
//let <identifer> = <expresiion>;
type Node interface {
	TokenLiteral() string
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
