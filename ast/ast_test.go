package ast

import (
	"github.com/stretchr/testify/assert"
	"monkey/token"
	"testing"
)

func Test_toString테스트(t *testing.T) {
	//given
	expectedInput := "let x = y;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
		},
	}

	//when
	result := program.ToString()

	//Then
	assert.Equal(t, expectedInput, result,
		"program.toString()이 다릅니다.\nexpected = %s, got = %q", expectedInput, result)
}
