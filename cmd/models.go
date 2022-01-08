package cmd

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jxsl13/animatch/anidb"
	"github.com/jxsl13/animatch/clean"
	"github.com/jxsl13/animatch/common"
	"github.com/jxsl13/animatch/filter"
	"github.com/spf13/cobra"
)

func Match(cmd *cobra.Command, paths []string) (MatchResults, error) {

	result := make(MatchResults, 0, len(paths))
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	cLimiter := make(chan struct{}, runtime.NumCPU())

	wg.Add(len(paths))
	for _, p := range paths {
		p := p
		cLimiter <- struct{}{}

		go func(path string, w *sync.WaitGroup) {
			defer w.Done()
			defer func() { <-cLimiter }()

			pathPrefix := clean.RemoveExtension(p)

			normalizedTerms := clean.LanguageTags(
				clean.ScreenResolutions(
					clean.TokenizeAll(
						clean.SplitPath(
							clean.Domains(
								clean.Tags(pathPrefix),
							), 1))))

			normalizedTerm := strings.Join(normalizedTerms, " ")

			distance, title, animeT, err := anidb.Search(normalizedTerm, filter.Metrics)
			if err != nil {
				common.Println(cmd, err)
				return
			}

			// lock for appending to result
			// also synchronize prinln
			mu.Lock()
			defer mu.Unlock()

			match := MatchResult{
				MatchTerm:  p,
				SearchTerm: normalizedTerm,
				FoundTitle: *title,
				FoundID:    animeT.AID,
				Distance:   *distance,
			}
			common.Println(cmd, match.String())

			result = append(result, match)
		}(p, wg)
	}

	wg.Wait()

	result.Sort()

	return result, nil
}

type MatchResult struct {
	MatchTerm  string
	SearchTerm string // normalized search term
	FoundTitle string
	FoundID    int
	Distance   float64
}

func (m *MatchResult) IsMatch(lowerBoundary float64) bool {
	return m.Distance >= lowerBoundary
}

func (m *MatchResult) Episode() string {
	if m.SearchTerm == "" {
		return ""
	}

	titleTokens := clean.Tokenize(m.FoundTitle)

	// only contains non title information. so integers in the title
	// should not disturb our search for the episode number
	lookupDiff := clean.RemoveEachOnce(m.SearchTerm, titleTokens...)

	// replace all information that is found in the title
	episode, found := filter.ExtractEpisode(lookupDiff)
	if !found {
		return ""
	}
	return episode

}

func (m *MatchResult) Tag() string {
	if m.FoundID == 0 {
		return ""
	}
	return fmt.Sprintf("[anidb-%d]", m.FoundID)
}

func (m *MatchResult) TaggedPath() string {
	if m.FoundID == 0 {
		return ""
	}

	pathPrefix := clean.RemoveExtension(m.MatchTerm)
	pathSuffix := filepath.Ext(m.MatchTerm)

	return pathPrefix + m.Tag() + pathSuffix
}

func (m *MatchResult) String() string {
	if m.FoundID == 0 {
		return ""
	}

	sb := strings.Builder{}
	sb.Grow(128 + len(m.MatchTerm) + len(m.SearchTerm) + len(m.FoundTitle))
	sb.WriteString("Match   : ")
	sb.WriteString(m.MatchTerm)
	sb.WriteString("\n")

	sb.WriteString("Search  : ")
	sb.WriteString(m.SearchTerm)
	sb.WriteString("\n")

	sb.WriteString("Episode : ")
	sb.WriteString(m.Episode())
	sb.WriteString("\n")

	sb.WriteString("Found   : ")
	sb.WriteString(m.FoundTitle)
	sb.WriteString("\n")

	sb.WriteString("ID      : ")
	sb.WriteString(strconv.Itoa(m.FoundID))
	sb.WriteString("\n")

	sb.WriteString("Distance: ")
	sb.WriteString(common.FormatFloat64(m.Distance))
	sb.WriteString("\n")
	sb.WriteString("\n")

	return sb.String()
}

type MatchResults []MatchResult

func (m MatchResults) Sort() {
	sort.Sort(byDistance(m))
}

func (m MatchResults) LongestMatchTerm() int {
	i := 0
	for _, t := range m {
		if i < len(t.MatchTerm) {
			i = len(t.MatchTerm)
		}
	}

	return i
}

func (m MatchResults) LongestTaggedPath() int {
	i := 0
	for _, t := range m {
		p := t.TaggedPath()
		if i < len(p) {
			i = len(p)
		}
	}

	return i
}

type byDistance []MatchResult

func (a byDistance) Len() int           { return len(a) }
func (a byDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDistance) Less(i, j int) bool { return a[i].Distance > a[j].Distance }
