package anidb

import (
	"math"

	"github.com/adrg/strutil"
	"github.com/jxsl13/animatch/clean"
	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"go.felesatra.moe/anidb"
)

const (
	ErrTitleNotFound = common.Error("title not found")

	// DefaultMatchDistanceUpperBound is the number of characters that
	// can at most be edited in order to reach the target search term
	// this is the default value that is set if no edit distance is provided as parameter
	DefaultMatchDistanceUpperBound = 5
)

type TitlesCache = anidb.TitlesCache
type StringMetric = strutil.StringMetric
type Anime = anidb.Anime
type AnimeT = anidb.AnimeT
type SearchResult []AnimeT

func (sr SearchResult) Titles() [][]string {
	titles := make([][]string, len(sr))

	for idx, a := range sr {
		titles[idx] = make([]string, 0, len(a.Titles))

		for _, t := range a.Titles {
			titles[idx] = append(titles[idx], t.Name)
		}
	}

	return titles
}

func MetaData(aid int) (*Anime, error) {
	client := NewClient()
	return client.RequestAnime(aid)
}

// Search simply looks for the provided terms
// maxEditDistance is an optional single value that allows to provide an upper boundary
// which is exclusive. This boundary controls the search match accuracy.
func Search(terms string, metric StringMetric) (*float64, *string, *AnimeT, error) {
	defaultCache, err := DefaultTitlesCache()
	if err != nil {
		return nil, nil, nil, err
	}
	defaultTitles, err := defaultCache.GetTitles()
	if err != nil {
		return nil, nil, nil, err
	}

	normalizedCache, err := NormalizedTitlesCache()
	if err != nil {
		return nil, nil, nil, err
	}

	normalizedTitles, err := normalizedCache.GetTitles()
	if err != nil {
		return nil, nil, nil, err
	}
	distance, animeIndex, titleIndex := search(terms, normalizedTitles, metric)
	// assuming that defaultCache index is the same as normalized cache index

	at := defaultTitles[animeIndex]
	title := at.Titles[titleIndex].Name

	return &distance, &title, &at, nil
}

func titlesToStrings(ts []anidb.Title) []string {
	result := make([]string, 0, len(ts))
	for _, t := range ts {
		result = append(result, t.Name)
	}
	return result
}

func search(terms string, ts []AnimeT, metric StringMetric) (distance float64, animeIndex, titleIndex int) {
	normalizedTerms := clean.Normalize(terms)

	maxDist := -math.MaxFloat64
	maxAnimeIndex := 0
	maxTitleIndex := 0

	for aIdx, t := range ts {
		// regex match
		titles := titlesToStrings(t.Titles)
		distance, tIdx, err := filter.BestMatch(normalizedTerms, titles, metric)
		if err != nil {
			continue
		}

		if distance > maxDist {
			maxDist = distance
			maxAnimeIndex = aIdx
			maxTitleIndex = tIdx
		}
	}
	return maxDist, maxAnimeIndex, maxTitleIndex
}
