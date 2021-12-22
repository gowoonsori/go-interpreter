package lexer

import (
	"monkey/token"
	"testing"
)

func Test기본연산자토큰테스트(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for i, testToken := range tests {
		token := lexer.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q but, got=%q",
				i, testToken.expectedType, token.Type)
		}
		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q but, got=%q",
				i, testToken.expectedLiteral, token.Literal)
		}
	}
}

func Test기본구문테스트(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = function(x, y){
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for i, testToken := range tests {
		token := lexer.NextToken()

		if token.Type != testToken.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q but, got=%q",
				i, testToken.expectedType, token.Type)
		}
		if token.Literal != testToken.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q but, got=%q",
				i, testToken.expectedLiteral, token.Literal)
		}
	}
}

func Test_추가연산자들테스트(t *testing.T) {
	input := `
	!-/*5
	5 < 10 > 5;

	if ( 5< 10) {
		return true;
	}else {
		return false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for i, testToken := range tests {
		tk := lexer.NextToken()

		if tk.Type != testToken.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q but, got=%q",
				i, testToken.expectedType, tk.Type)
		}
		if tk.Literal != testToken.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q but, got=%q",
				i, testToken.expectedLiteral, tk.Literal)
		}
	}
}

func Test_두문자토큰테스트(t *testing.T) {
	input := `10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for i, testToken := range tests {
		tk := lexer.NextToken()

		if tk.Type != testToken.expectedType {
			t.Fatalf("tests [%d] - tokentype wrong. expected=%q but, got=%q",
				i, testToken.expectedType, tk.Type)
		}
		if tk.Literal != testToken.expectedLiteral {
			t.Fatalf("tests [%d] - literal wrong. expected=%q but, got=%q",
				i, testToken.expectedLiteral, tk.Literal)
		}
	}
}
