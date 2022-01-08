package clean

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Normalize(s string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))

	// DefaultNormalizationTransformer is used to transform stringt before
	// computing their distances
	// Reference: https://go.dev/blog/normalization
	transformer := transform.Chain(norm.NFC, norm.NFKC)
	io.Copy(buf, transform.NewReader(strings.NewReader(s), transformer))
	return buf.String()
}

func NormalizeAll(ss []string) []string {
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		result = append(result, Normalize(s))
	}
	return result
}
