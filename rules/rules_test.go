package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHints(t *testing.T) {
	tern := HintFromColors("ggxyx")
	assert.Equal(t, "ggxyx", tern.ToColors())

	tern = HintFromColors("xxxxx")
	assert.Equal(t, "xxxxx", tern.ToColors())

	tern = HintFromColors("yyyyy")
	assert.Equal(t, "yyyyy", tern.ToColors())

	tern = HintFromColors("ggggg")
	assert.Equal(t, "ggggg", tern.ToColors())

	tern = HintFromColors("xxyyg")
	assert.Equal(t, "xxyyg", tern.ToColors())

	tern = HintFromColors("yyxgg")
	assert.Equal(t, "yyxgg", tern.ToColors())
}

func TestCheck(t *testing.T) {
	tests := []struct {
		guess  string
		secret string
		hint   string
	}{
		{guess: "yummy", secret: "mummy", hint: "xgggg"},
		{guess: "slate", secret: "dream", hint: "xxyxy"},
		{guess: "cigar", secret: "vodka", hint: "xxxyx"},
		{guess: "pound", secret: "fanny", hint: "xxxgx"},
		{guess: "nanny", secret: "fanny", hint: "xgggg"},
		{guess: "error", secret: "rebut", hint: "yyxxx"},
	}

	for _, test := range tests {
		assert.Equal(t, test.hint, Check(test.guess, test.secret))
	}
}
