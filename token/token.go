package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//식별자 + 리터럴
	IDENT = "IDENT" //add, foobar, x, y ...
	INT   = "INT"   //123456

	//연산자
	ASSIGN = "="
	PLUS   = "+"

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
)

//예약 키워드
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
