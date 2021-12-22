package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

//구문에서 다음 token 셋팅
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram AST로 파싱 하고 루트노드 리턴
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

//Errors  Parser의 에러들 반환
func (p *Parser) Errors() []string {
	return p.errors
}

//문장 시작토큰 보고 적절한 구문파서로 파싱
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

//let 절 구문 파싱
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	//항상 변수명으로 시작하기 때문에 검사
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	//현재 변수명을 식별자로 인식
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//대입 연산자 검사
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//구문 끝은 세미콜론이기 때문에 세미콜론이 올때 까지 계속 읽어 값으로 인식
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
		//input이 끝났는데도 세미콜론이 없으면 에러
		if p.curToken.Type == token.EOF {
			p.addError(token.SEMICOLON)
			return nil
		}
	}

	return stmt
}

//파서의 현재 토큰이 parameter와 같은지 check
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

//현재 파싱할 다음 토큰이 parmeter 인지 check하고 맞으면 토큰 이동
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addError(t)
		return false
	}
}

//현재 peek 토큰이 parameter와 같은지 check
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) addError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
