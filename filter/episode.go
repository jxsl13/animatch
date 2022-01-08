package filter

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	// 0 - 1950 is themax number of allowed episodes (e.g. One PIece has got like 1050 episodes as of now, so we got another 20 years to go before we hit potential years (1950 might be a year but I doubt there were anime that anyoe is matching against as of now.))
	episodeNumberRegex = regexp.MustCompile(`\b(0*[0-9]|[1-8][0-9]|9[0-9]|[1-8][0-9]{2}|9[0-8][0-9]|99[0-9]|1[0-8][0-9]{2}|19[0-4][0-9]|1950)\b`)
)

// ExtractEpisode tries to extract a numeric value.
// Returns the string match and whether something was found
func ExtractEpisode(searchTerm string) (string, bool) {

	// episode number should be in the second half of the string
	ss := episodeNumberRegex.FindAllString(searchTerm, -1)
	if ss == nil {

		return "", false
	}

	// wherever we found some matches, we treat them the same way at this point
	if len(ss) == 1 {
		// remove leading zeroes
		return strings.TrimLeft(ss[0], "0"), true
	}

	numbers := make([]int, 0, len(ss))
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		numbers = append(numbers, i)
	}

	sort.Ints(numbers)
	return strconv.Itoa(numbers[0]), true
}
