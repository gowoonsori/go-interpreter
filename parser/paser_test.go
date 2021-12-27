package parser

import (
	"fmt"
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

func Test_PrefixOperater_Expression(t *testing.T) {
	//given
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	//when
	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()

		//then
		checkParserErrors(t, p)

		assert.Equalf(t, 1, len(program.Statements), "statement가 1개가 아닙니다. got = %d", len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.Equalf(t, true, ok, "표현식이 아닙니다. got = %T", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.Equalf(t, true, ok, "전위 표현식이 아닙니다. got = %T", stmt.Expression)
		assert.Equalf(t, tt.operator, exp.Operator, "표현식의 전위 연산자가 %s가 아닙니다. got = %s", tt.operator, exp.Operator)
		testIntegerLiteral(t, exp.Right, tt.integerValue)
	}
}

func Test_InfixOperator_Expression(t *testing.T) {
	//given
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"5 / 5;", 5, "/", 5},
	}

	//when
	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()

		//then
		checkParserErrors(t, p)

		assert.Equalf(t, 1, len(program.Statements), "statement가 1개가 아닙니다. got = %d", len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.Equalf(t, true, ok, "statement가 표현식이 아닙니다. got = %T", program.Statements[0])

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.Equalf(t, true, ok, "중위 표현식이 아닙니다. got = %T", stmt.Expression)

		testIntegerLiteral(t, exp.Left, tt.leftValue)
		assert.Equalf(t, tt.operator, exp.Operator, "연산자가 %s와 다릅니다. got = %s", tt.operator, exp.Operator)
		testIntegerLiteral(t, exp.Right, tt.rightValue)
	}
}

func Test_OperatorPrecedence(t *testing.T) {
	//given
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a+b*c+d/e-f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5* 5", "(3 + 4)((-5) * 5)"},
		{"5 >4 == 3 >4", "((5 > 4) == (3 > 4))"},
		{"5 < 4 != 3>4", "((5 < 4) != (3 > 4))"},
		{"3 + 4*5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
	}

	//when
	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()

		//then
		checkParserErrors(t, p)

		actual := program.ToString()
		assert.Equalf(t, tt.expected, actual, "expected = %q, got = %q", tt.expected, actual)
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, expectedValue int64) {
	integ, ok := il.(*ast.IntegerLiteral)
	assert.Equalf(t, true, ok, "정수 리터럴이 아닙니다. got = %T", il)
	assert.Equalf(t, expectedValue, integ.Value, "정수 값이 %d와 다릅니다. got = %d", expectedValue, integ.Value)
	assert.Equalf(t, fmt.Sprintf("%d", expectedValue), integ.TokenLiteral(),
		"토큰 리터럴이 %d와 다릅니다. got = %s", expectedValue, integ.TokenLiteral())
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
