package clean

import (
	"os"
	"strings"

	"github.com/jxsl13/animatch/common"
)

func RemoveExtension(filePath string) string {

	if pos := strings.LastIndexByte(filePath, '.'); pos != -1 {
		return filePath[:pos]
	}
	return filePath
}

func SplitPath(p string, reverseDepth ...int) []string {
	n := common.OptionalIntWithDefault(-1, reverseDepth...)
	parts := strings.Split(p, string(os.PathSeparator))
	parts = RemoveEmpty(parts)
	if n <= 0 || len(parts) <= n {
		return parts
	}

	return parts[len(parts)-n:]
}

func Overlap(tokens []string) []string {
	result := make([]string, 0, len(tokens))

outer:
	for _, above := range tokens {
		if len(result) == 0 {
			// add first
			result = append(result, above)
			continue
		}

		for _, below := range result {
			if below == above {
				continue outer
			}
		}
		result = append(result, above)
	}
	return result
}
