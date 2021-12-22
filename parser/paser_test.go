package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func Test_Parser테스트(t *testing.T) {
	//given
	input := `
let x = 5;
let y = 10;
let z = 838383;
`
	//when
	l := lexer.NewLexer(input)
	p := New(l)
	program := p.ParseProgram()

	//then
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("3개의 구문이 아닙니다. expected : 3, got: %d", len(program.Statements))
	}
	checkParserErrors(t, p)
}

func Test_Parser_실패테스트(t *testing.T) {
	//given
	input := `let foo = 8`

	//when
	l := lexer.NewLexer(input)
	p := New(l)
	p.ParseProgram()

	//then
	if len(p.Errors()) != 1 {
		t.Fatalf("error가 한개이어야 합니다.")
	}
}

func Test_let구문테스트(t *testing.T) {
	//given
	input := `
let x x 5;
let y = 10;
let foo = 838383;
`
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	//when
	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()

	//then
	checkParserErrors(t, p)

	for i, tk := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tk.expectedIdentifier) {
			return
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

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral 이 let이 아닙니다. got=%q ", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s 가 LetStatemnt가 아닙니다. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("name이 %s가 아닙니다. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("token literal이 %s가 아닙니다. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
