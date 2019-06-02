package word

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("mmoomm") {
		t.Error(`IsPalindrome("hello")=false`)
	}
}
