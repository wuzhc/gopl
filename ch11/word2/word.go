// 对比word1版本,使用字符序列比较也不是字节序列比较,忽略非字母,大小写
package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
