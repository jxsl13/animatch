package filter

import (
	"math"

	"github.com/adrg/strutil"
)

func BestMatch(comp string, selection []string, metric strutil.StringMetric) (distance float64, index int, err error) {
	if len(selection) == 0 {
		return -1, -1, ErrSelectionEmpty
	}

	maxIdx := -1
	maxDistance := -math.MaxFloat64
	for idx, s := range selection {
		distance := metric.Compare(s, comp)
		if distance > maxDistance {
			maxDistance = distance
			maxIdx = idx
		}
	}

	return maxDistance, maxIdx, nil
}

func BestMatchNested(comp string, selections [][]string, metric strutil.StringMetric) (distance float64, outerIndex, innerIndex int, err error) {
	if len(selections) == 0 {
		return 0, -1, -1, ErrSelectionEmpty
	}

	emptySelections := 0
	maxOuterIndex := -1
	maxInnerIndex := -1
	maxDistance := -math.MaxFloat64

	for oIdx, selection := range selections {
		if len(selection) == 0 {
			emptySelections++
			continue
		}

		distance, idx, _ := BestMatch(comp, selection, metric)
		if distance > maxDistance {
			maxDistance = distance
			maxOuterIndex = oIdx
			maxInnerIndex = idx
		}

	}

	if len(selections) == emptySelections {
		return 0, -1, -1, ErrSelectionEmpty
	}
	return maxDistance, maxOuterIndex, maxInnerIndex, nil
}
