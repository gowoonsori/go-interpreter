package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, ch string) *Token {
	return &Token{Type: tokenType, Literal: ch}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//식별자 + 리터럴
	IDENT = "IDENT" //add, foobar, x, y ...
	INT   = "INT"   //123456

	//연산자
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	LT = "<"
	GT = ">"

	//구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

//예약 키워드
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"true":     TRUE,
	"false":    FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
