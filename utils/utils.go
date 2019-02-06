package utils

import (
	"regexp"
	"strings"
)

var splitRegexp = regexp.MustCompile("\\s+")

func Min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func GetFirstNWords(s string, N int) string {
	wordsSlice := splitRegexp.Split(s, N+1)
	count := Min(N, len(wordsSlice))
	return strings.Join(wordsSlice[:count], " ")
}
