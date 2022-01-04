package clean

import "golang.org/x/text/language"

var (
	languageTagList = []string{
		"subbed",
		"dubbed",
	}

	LanguageTagMap = map[string]bool{}
)

func init() {
	for _, tag := range languageTagList {
		LanguageTagMap[tag] = true
	}
}

func Language(ss []string) []string {
	result := make([]string, 0, len(ss))

	for _, s := range ss {
		if _, err := language.ParseBase(s); err != nil {
			result = append(result, s)
		}
	}
	return result
}

func LanguageTags(ss []string) []string {
	result := make([]string, 0, len(ss))

	for _, s := range ss {
		if !LanguageTagMap[s] {
			result = append(result, s)
		}
	}
	return result
}
