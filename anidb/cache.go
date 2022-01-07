package anidb

import (
	"strings"
	"sync"

	"github.com/jxsl13/animatch/clean"
	"go.felesatra.moe/anidb"
)

var (
	DecodeTitles = anidb.DecodeTitles

	onceNormalizedCache = sync.Once{}
	normalizedCache     = (*TitlesCache)(nil)

	onceDefaultTitles = sync.Once{}
	defaultCache      = (*TitlesCache)(nil)
)

// DefaultTitlesCache returns a singleton of the solely once parsed titles cache that are stored locally.
func DefaultTitlesCache() (*TitlesCache, error) {
	if defaultCache != nil {
		return defaultCache, nil
	}
	var err error
	onceDefaultTitles.Do(func() {
		defaultCache, err = anidb.DefaultTitlesCache()
	})

	if err != nil {
		return nil, err
	}
	return defaultCache, nil
}

// NormalizedTitlesCache returns a singleton where every title is normalized.
// This instance is used in order to speed up the search algorithms that redundantly
// normalize this cache's titles
func NormalizedTitlesCache() (*TitlesCache, error) {
	if normalizedCache != nil {
		return normalizedCache, nil
	}
	var err error
	onceNormalizedCache.Do(func() {

		// get a fresh copy of the cache
		var tc *TitlesCache
		tc, err = anidb.DefaultTitlesCache()
		if err != nil {
			return
		}

		// start num cpu worker routines
		for i, anime := range tc.Titles {

			// normalize all titles
			for j, title := range anime.Titles {
				name := title.Name
				normalizedName := strings.Join(clean.NormalizeAll(clean.Tokenize(name)), " ")
				tc.Titles[i].Titles[j].Name = normalizedName
			}

		}

		normalizedCache = tc
	})

	if err != nil {
		return nil, err
	}
	return normalizedCache, nil
}
