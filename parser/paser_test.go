package parser

import (
	"github.com/stretchr/testify/assert"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func Test_Let_statement(t *testing.T) {
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

func Test_Return_statement(t *testing.T) {
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

func Test_Expression_구문테스트(t *testing.T) {
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

func Test_Identifier_표현식_테스트(t *testing.T) {
	//given
	input := "foobar;"

	//when
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	//then
	checkParserErrors(t, p)

	assert.Equalf(t, 1, len(program.Statements),
		"program 의 구문이 1개가 아닙니다. got = %d ", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.Equalf(t, true, ok, "표현식이 아닙니다. got = %T", program.Statements[0])

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.Equalf(t, true, ok, "Identifier 이 아닙니다. got = %T", stmt.Expression)
	assert.Equalf(t, "foobar", ident.Value, "ident의 값이 %s 가 아닙니다. got = %s", "foobar", ident.Value)
	assert.Equalf(t, "foobar", ident.TokenLiteral(), "ident의 값이 %s 가 아닙니다. got = %s", "foobar", ident.TokenLiteral())
}

func Test_IntegerLiteral_표현식_테스트(t *testing.T) {
	//given
	input := "5;"
	l := lexer.NewLexer(input)
	p := NewParser(l)

	//when
	program := p.ParseProgram()

	//then
	checkParserErrors(t, p)
	assert.Equalf(t, 1, len(program.Statements),
		"program 의 구문이 1개가 아닙니다. got = %d ", len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.Equalf(t, true, ok, "표현식이 아닙니다. got = %T", program.Statements[0])

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.Equalf(t, true, ok, "IntegerLiteral 이 아닙니다. got = %T", stmt.Expression)
	assert.Equalf(t, int64(5), ident.Value, "ident의 값이 %d 가 아닙니다. got = %d", 5, ident.Value)
	assert.Equalf(t, "5", ident.TokenLiteral(), "ident의 값이 %s 가 아닙니다. got = %s", "5", ident.TokenLiteral())
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
