package word

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
