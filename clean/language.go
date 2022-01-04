package clean

import "golang.org/x/text/language"

func Language(ss []string) []string {
	result := make([]string, 0, len(ss))

	for _, s := range ss {
		if _, err := language.ParseBase(s); err != nil {
			result = append(result, s)
		}
	}
	return result
}
