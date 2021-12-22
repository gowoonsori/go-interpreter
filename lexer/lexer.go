package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

const EOF = 0

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = *token.New(token.EQ, literal)
		} else {
			tok = *token.New(token.ASSIGN, string(l.ch))
		}
	case ';':
		tok = *token.New(token.SEMICOLON, string(l.ch))
	case '(':
		tok = *token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = *token.New(token.RPAREN, string(l.ch))
	case '{':
		tok = *token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = *token.New(token.RBRACE, string(l.ch))
	case ',':
		tok = *token.New(token.COMMA, string(l.ch))
	case '+':
		tok = *token.New(token.PLUS, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = *token.New(token.NOT_EQ, literal)
		} else {
			tok = *token.New(token.BANG, string(l.ch))
		}
	case '-':
		tok = *token.New(token.MINUS, string(l.ch))
	case '*':
		tok = *token.New(token.ASTERISK, string(l.ch))
	case '/':
		tok = *token.New(token.SLASH, string(l.ch))
	case '<':
		tok = *token.New(token.LT, string(l.ch))
	case '>':
		tok = *token.New(token.GT, string(l.ch))
	case EOF:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = *token.New(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return EOF
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
