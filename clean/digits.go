package clean

import "regexp"

var (
	digitReplacer = regexp.MustCompile(`\d+`)
)

func NoDigits(tokens []string) []string {
	result := make([]string, 0, len(tokens))

	for _, token := range tokens {
		result = append(result, SplitSpace(ShrinkSpace(digitReplacer.ReplaceAllString(token, " ")))...)
	}
	return result
}
