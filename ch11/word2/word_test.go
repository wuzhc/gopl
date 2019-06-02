package word

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"a man nam a", true},
		{"aba", true},
		{"hello", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("isPalindrome(%q)=%v", test.input, got)
		}
	}
}
