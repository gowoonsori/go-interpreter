package parser

import (
	"github.com/stretchr/testify/assert"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func Test_let구문테스트(t *testing.T) {
	//given
	input := `
		let x = 5;
		let y = 10;
		let foo = 838383;`
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	//when
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()

	//then
	assert.Equalf(t, 3, len(program.Statements), "3개의 구문이 아닙니다. expected : 3, got: %d", len(program.Statements))
	checkParserErrors(t, p)

	for i, tk := range tests {
		stmt := program.Statements[i]
		assert.Equalf(t, "let", stmt.TokenLiteral(), "s.TokenLiteral 이 let이 아닙니다. got=%q ", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		assert.Equalf(t, true, ok, "s 가 LetStatemnt가 아닙니다. got=%T", stmt)
		assert.Equalf(t, tk.expectedIdentifier, letStmt.Name.Value,
			"name이 %s가 아닙니다. got=%s", tk.expectedIdentifier, letStmt.Name.Value)
		assert.Equalf(t, tk.expectedIdentifier, letStmt.Name.TokenLiteral(),
			"token literal이 %s가 아닙니다. got=%s", tk.expectedIdentifier, letStmt.Name.TokenLiteral())
	}
}

func Test_return구문테스트(t *testing.T) {
	//given
	input := `
		return 5;
		return 10;
		return 993232;`

	//when
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	//then
	assert.Equalf(t, 3, len(program.Statements),
		"3개의 구문이 아닙니다. expected : 3, got: %d", len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.Equalf(t, true, ok, "Stmt Return구문이 아닙니다. got = %T", stmt)
		assert.Equalf(t, "return", returnStmt.TokenLiteral(),
			"Return 구문의 토큰 리터럴이 'return'이 아닙니다. got = %q", returnStmt.TokenLiteral())
	}
}

func Test_expression구문테스트(t *testing.T) {
	//given
	input := `
		return 5;
		return 10;
		return 993232;`

	//when
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("3개의 구문이 아닙니다. expected : 3, got: %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Stmt Return구문이 아닙니다. got = %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Return 구문의 토큰 리터럴이 'return'이 아닙니다. got = %q", returnStmt.TokenLiteral())
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}