package clean

var (
	resolutionList = []string{
		"720p",
		"1080p",
	}

	ResolutionMap = map[string]bool{}
)

func init() {
	for _, res := range resolutionList {
		ResolutionMap[res] = true
	}
}

func Resolutions(ss []string) []string {
	result := make([]string, 0, len(ss))

	for _, s := range ss {
		// skip matches
		if !ResolutionMap[s] {
			result = append(result, s)
		}
	}
	return result
}
