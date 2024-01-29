package palindromo

import (
	"strings"
)

func EsPalindromo(s string) bool {
	return esPalindromo(strings.ToLower(s), 0, len(s)-1)
}

func esPalindromo(s string, i int, f int) bool {
	if f-i < 2 {
		return true
	}
	if s[i] == s[f] {
		return esPalindromo(s, i+1, f-1)
	}
	return false
}
