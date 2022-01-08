package clean

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/adrg/strutil"
	"github.com/jxsl13/animatch/common"
)

func UniqueSlice(ss []string) []string {
	return strutil.UniqueSlice(ss)
}

func RemoveEmpty(ss []string) []string {
	result := make([]string, 0, len(ss))

	for _, s := range ss {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}

func ToLowerCase(ss []string) []string {
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		result = append(result, strings.ToLower(s))
	}

	return result
}

func TokenizeEach(ss []string, delimiter ...string) []string {
	result := make([]string, 0, len(ss))
	delim := common.OptionalStringWithDefault(" ", delimiter...)

	for _, t := range ss {
		if normalized := Tokenize(t); len(normalized) > 0 {
			result = append(result, strings.Join(normalized, delim))
		}
	}
	return result
}

func TokenizeAll(terms []string) []string {
	result := make([]string, 0, len(terms))

	for _, t := range terms {
		if normalized := Tokenize(t); len(normalized) > 0 {
			result = append(result, normalized...)
		}
	}

	return result

}

func Tokenize(token string) []string {
	tokens := SplitSpace(ShrinkSpace(ReplaceSpecial(strings.ToLower(strings.TrimSpace(token)))))

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

func RemoveEachOnce(s string, removeSet ...string) string {

	for _, rem := range removeSet {
		s = strings.Replace(s, rem, "", 1)
	}

	return s
}
