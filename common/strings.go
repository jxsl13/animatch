package common

import "strings"

func ToLowerCaseAll(ss []string) []string {
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		result = append(result, strings.ToLower(s))
	}
	return result
}
