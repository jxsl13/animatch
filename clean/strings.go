package clean

import "strings"

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
