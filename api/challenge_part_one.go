package api

import (
	"strings"
	"unicode"
)

// Solution for part one of coding challenge https://gist.github.com/charliemcelfresh/ffd8855600fa834b28f1abb29f0d417d
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	var ch string
	Result := ""
	idx := 0
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			// tracking only indexes for alphanumeric
			idx++
			if idx%3 == 0 {
				ch = strings.ToUpper(string(c))
			} else {
				ch = strings.ToLower(string(c))
			}
			Result += ch
		} else {
			Result += string(c)
		}
	}
	return Result
}
