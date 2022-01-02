package clean

import (
	"regexp"
	"strings"
	"unicode"
)

func NormalizeAll(terms []string) []string {
	result := make([]string, 0, len(terms))

	for _, t := range terms {
		if normalized := Normalize(t); len(normalized) > 0 {
			result = append(result, normalized...)
		}
	}

	return result

}

func Normalize(token string) []string {
	tokens := SplitSpace(ShrinkSpace(ReplaceSpecial(strings.TrimSpace(strings.ToLower(token)))))

	return tokens
}

func isAnyOf(r rune, fs ...func(r rune) bool) bool {
	for _, f := range fs {
		if f(r) {
			return true
		}
	}
	return false
}

func ReplaceSpecial(token string) string {
	result := make([]rune, 0, len(token)/4)
	for _, r := range token {
		if isAnyOf(r, unicode.IsMark, unicode.IsPunct, unicode.IsSymbol) {
			r = ' '
		}
		result = append(result, r)
	}

	return string(result)
}

var (
	spaceReplacer = regexp.MustCompile(`\s+`)
)

func ShrinkSpace(token string) string {
	return spaceReplacer.ReplaceAllString(token, " ")
}

func SplitSpace(token string) []string {
	split := spaceReplacer.Split(token, -1)
	result := make([]string, 0, len(split))

	for _, s := range split {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
