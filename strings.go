package crt

import (
	"strings"

	lang "github.com/mt1976/crt/language"
	styl "github.com/mt1976/crt/styles"
)

// The upcase function in Go converts a string to uppercase.
// upcase converts a string to uppercase.
func upcase(s string) string {
	return strings.ToUpper(s)
}

func downcase(s string) string {
	return strings.ToLower(s)
}

// The function `trimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
// a new string with all consecutive occurrences of `c` trimmed down to a single occurrence.
func trimRepeatingCharacters(s string, c string) string {

	result := ""
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		if i == 0 {
			result = string(s[i])
		} else {
			if string(s[i]) != c || string(s[i-1]) != c {
				result = result + string(s[i])
			}
		}
	}
	return result
}

func bold(s string) string {
	return styl.Bold + s + styl.Reset
}

func sQuote(s string) string {
	return lang.SymSingleQuote + s + lang.SymSingleQuote
}

func pQuote(s string) string {
	return lang.SymOpenBracket + s + lang.SymCloseBracket
}

func dQuote(s string) string {
	return lang.SymDoubleQuote + s + lang.SymDoubleQuote
}

func qQuote(s string) string {
	return lang.SymSquareQuoteOpen + s + lang.SymSquareQuoteClose
}

func isActionIn(in string, check ...string) bool {
	for i := 0; i < len(check); i++ {
		if strings.Contains(upcase(in), check[i]) {
			return true
		}
	}
	return false
}
