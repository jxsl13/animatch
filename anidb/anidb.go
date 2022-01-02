package anidb

import (
	"regexp"
	"strings"

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

var (
	wrapErrTitleNotFound = common.WrapErr(ErrTitleNotFound)
)

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

// BestMatch tries to fetch the best matching anime for the provided search term
// maxEditDistance is an optional single value that allows to provide an upper boundary
// which is exclusive. This boundary controls the search match accuracy.
func BestMatch(terms []string, maxEditDistance ...int) (title string, at *AnimeT, err error) {
	ts, err := Search(terms, maxEditDistance...)
	if err != nil {
		return "", nil, err
	}

	titles := ts.Titles()
	_, i, j, err := filter.LowestLDistNestedArgs(terms, titles)
	if err != nil {
		return "", nil, err
	}

	bestTitle := titles[i][j]
	bestAnime := ts[i]
	return bestTitle, &bestAnime, nil
}

// Search simply looks for the provided terms
// maxEditDistance is an optional single value that allows to provide an upper boundary
// which is exclusive. This boundary controls the search match accuracy.
func Search(terms []string, maxEditDistance ...int) (SearchResult, error) {
	tc, err := anidb.DefaultTitlesCache()
	if err != nil {
		return nil, err
	}
	defer tc.SaveIfUpdated()
	ts, err := tc.GetTitles()
	if err != nil {
		return nil, err
	}
	ts = search(terms, ts, maxEditDistance...)
	if len(ts) == 0 {
		return nil, wrapErrTitleNotFound(strings.Join(terms, " "))
	}
	return ts, nil
}

// globTerms returns a regexp that matches strings containing the
// terms in order, ignoring case and intervening characters.
func globTerms(terms []string) *regexp.Regexp {
	for i, t := range terms {
		terms[i] = regexp.QuoteMeta(t)
	}

	normalizedLcTerms := filter.NormalizeAll(common.ToLowerCaseAll(terms))

	return regexp.MustCompile("(?i)" + strings.Join(normalizedLcTerms, ".*"))
}

func titlesToStrings(ts []anidb.Title) []string {
	result := make([]string, 0, len(ts))
	for _, t := range ts {
		result = append(result, t.Name)
	}
	return result
}

// search returns a slice of anime whose title matches the given
// terms.  A title is matched if it contains all terms in order,
// ignoring case and intervening characters.
// If the regex is not matched, the best match is returned that has the lowest
// levenshtein edit distance that is also capped at an upper boundary that is concidered not
// a proper match anymore.
// maxEditDistance can be passed as single optional parameter that is used to vary the search accuracy
func search(terms []string, ts SearchResult, maxEditDistance ...int) SearchResult {
	r := globTerms(terms)
	term := strings.Join(terms, " ")

	// remove empty spaces that we addd by joining the strings
	// and allow some more leeway
	upperBound := (len(terms) - 1) + common.OptionalIntWithDefault(DefaultMatchDistanceUpperBound, maxEditDistance...)

	var matched SearchResult
	for _, at := range ts {
		// regex match
		titles := titlesToStrings(at.Titles)

		if titleMatches(r, titles) {
			matched = append(matched, at)
		} else if _, found := filter.DistanceBelowAny(upperBound, term, titles); found {
			matched = append(matched, at)
		}
	}
	return matched
}

// titleMatches returns true if any of the titles matches the regexp.
func titleMatches(r *regexp.Regexp, ts []string) bool {
	for _, t := range ts {
		if r.FindStringIndex(t) != nil {
			return true
		}
	}
	return false
}
