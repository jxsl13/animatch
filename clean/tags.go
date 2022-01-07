package clean

import "regexp"

var (
	tagReplacer = regexp.MustCompile(`\[[^\[\]]*\]`)
)

// ags removes all [...] tags from the string
func Tags(s string) string {

	return tagReplacer.ReplaceAllString(s, "")
}
