package clean

import (
	"regexp"
)

var (
	DomainRegexList = []*regexp.Regexp{
		regexp.MustCompile(`(?i)www\.[a-z0-9-]{1,30}\.[a-z]{2,}`),
	}
)

func Domains(s string) string {
	for _, re := range DomainRegexList {
		s = re.ReplaceAllString(s, "")
	}
	return s
}
