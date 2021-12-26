package main

import (
	"testing"
)

func TestTokenType(t *testing.T) {
	tests := []struct {
		name  string
		token TokenType
		want  string
	}{
		{
			name:  "left paren",
			token: LEFT_PAREN,
			want:  "LEFT_PAREN",
		},
		{
			name:  "and",
			token: AND,
			want:  "AND",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.want != test.token.String() {
				t.Errorf("Expect: %s, Got: %s", test.want, test.token)
			}
		})
	}
}
