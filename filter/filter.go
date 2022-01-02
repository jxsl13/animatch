package filter

import (
	"bytes"
	"io"
	"math"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/jxsl13/animatch/common"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	ErrSelectionEmpty = common.Error("selection is empty")
)

var (
	// DefaultNormalizationTransformer is used to transform stringt before
	// computing their distances
	// Reference: https://go.dev/blog/normalization
	DefaultNormalizationTransformer = transform.Chain(norm.NFC, norm.NFKC)
)

func Normalize(s string) string {
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	io.Copy(buf, transform.NewReader(strings.NewReader(s), DefaultNormalizationTransformer))
	return buf.String()
}

func NormalizeAll(ss []string) []string {
	result := make([]string, 0, len(ss))
	for _, s := range ss {
		result = append(result, Normalize(s))
	}
	return result
}

func Distance(s1, s2 string) int {
	return levenshtein.ComputeDistance(Normalize(s1), Normalize(s2))
}

func DistanceBelowSelect(upperBound int, s string, selection []string) (indices []int) {
	indices = make([]int, 0, 1)
	for i, sel := range selection {
		if DistanceBelow(upperBound, s, sel) {
			indices = append(indices, i)
		}
	}
	return indices
}

func DistanceBelow(upperBound int, s, s2 string) bool {
	return Distance(s, s2) < upperBound
}

func DistanceBelowAny(upperBound int, s string, selection []string) (int, bool) {
	for i, sel := range selection {
		if DistanceBelow(upperBound, s, sel) {
			return i, true
		}
	}
	return -1, false
}

func SliceFromIndex(ss []string, is []int) []string {
	l := len(ss)
	result := make([]string, 0, len(is))
	for _, i := range is {
		if i < 0 || l <= i {
			continue
		}
		result = append(result, ss[i])
	}
	return result
}

func LowestLDistArgs(args, selection []string) (distance, index int, err error) {
	return LowestLDist(strings.Join(args, " "), selection)
}

func LowestLDist(comp string, selection []string) (distance, index int, err error) {
	if len(selection) == 0 {
		return -1, -1, ErrSelectionEmpty
	}

	minIdx := -1
	minDistance := math.MaxInt
	for idx, s := range selection {
		distance := Distance(comp, s)
		if distance < minDistance {
			minDistance = distance
			minIdx = idx
		}
	}

	return minDistance, minIdx, nil
}

func LowestLDistNestedArgs(args []string, selections [][]string) (distance, outerIndex, innerIndex int, err error) {
	return LowestLDistNested(strings.Join(args, " "), selections)
}

func LowestLDistNested(comp string, selections [][]string) (distance, outerIndex, innerIndex int, err error) {
	if len(selections) == 0 {
		return 0, -1, -1, ErrSelectionEmpty
	}

	emptySelections := 0
	minOuterIndex := -1
	minInnerIndex := -1
	minDistance := math.MaxInt

	for oIdx, selection := range selections {
		if len(selection) == 0 {
			emptySelections++
			continue
		}

		distance, idx, _ := LowestLDist(comp, selection)
		if distance < minDistance {
			minDistance = distance
			minOuterIndex = oIdx
			minInnerIndex = idx
		}

	}

	if len(selections) == emptySelections {
		return 0, -1, -1, ErrSelectionEmpty
	}
	return minDistance, minOuterIndex, minInnerIndex, nil
}
