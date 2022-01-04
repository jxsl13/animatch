package filter

import (
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/jxsl13/animatch/common"
)

var (
	ErrSelectionEmpty = common.Error("selection is empty")

	MetricLevenshtein  = metrics.NewLevenshtein()
	MetricHamming      = metrics.NewHamming()
	MetricJaccard      = metrics.NewJaccard()
	MetricJaro         = metrics.NewJaro()
	MetricJaroWinkler  = metrics.NewJaroWinkler()
	MetricSorensenDice = metrics.NewSorensenDice()

	Metrics = ComposedStringMetric{
		MetricLevenshtein,
		MetricHamming,
		MetricJaccard,
		MetricJaro,
		MetricJaroWinkler,
		MetricSorensenDice,
	}

	MetricNames = []string{
		"Levenshtein",
		"Hamming",
		"Jaccard",
		"Jaro",
		"JaroWinkler",
		"SorensenDice",
	}
)

func init() {
	MetricLevenshtein.DeleteCost = 2
	MetricLevenshtein.InsertCost = 2
	MetricLevenshtein.ReplaceCost = 3

	MetricJaccard.NgramSize = 5

	MetricSorensenDice.NgramSize = 5
}

type ComposedStringMetric []strutil.StringMetric

func (cm ComposedStringMetric) Compare(a, b string) float64 {
	sum := 0.0

	for _, m := range cm {
		sum += m.Compare(a, b)
	}

	return sum / float64(len(cm))
}
