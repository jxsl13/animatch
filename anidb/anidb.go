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
)

var (
	// MatchDistanceUpperBound is th enumber of characters that
	// can at most be edited in order to reach the target search term
	MatchDistanceUpperBound = 5
	wrapErrTitleNotFound    = common.WrapErr(ErrTitleNotFound)
)

type SearchResult []anidb.AnimeT

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

// BestMatch tries to fetch the best matching anime for the provided search term
func BestMatch(terms []string) (title string, at *anidb.AnimeT, err error) {
	ts, err := Search(terms)
	if err != nil {
		return "", nil, err
	}

	titles := ts.Titles()
	_, i, j, err := filter.LowestLDistNestedArgs(terms, ts.Titles())
	if err != nil {
		return "", nil, err
	}

	bestTitle := titles[i][j]
	bestAnime := ts[i]
	return bestTitle, &bestAnime, nil
}

// Search simply looks for the provided terms
func Search(terms []string) (SearchResult, error) {
	tc, err := anidb.DefaultTitlesCache()
	if err != nil {
		return nil, err
	}
	defer tc.SaveIfUpdated()
	ts, err := tc.GetTitles()
	if err != nil {
		return nil, err
	}
	ts = search(terms, ts)
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
func search(terms []string, ts SearchResult) SearchResult {
	r := globTerms(terms)
	term := strings.Join(terms, " ")

	// remove empty spaces that we addd by joining the strings
	// and allow some more leeway
	upperBound := (len(terms) - 1) + MatchDistanceUpperBound

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
