package clean

var (
	resolutionList = []string{
		"360p",
		"480p",
		"720p",
		"1080p",
		"2160",
	}

	ResolutionMap = map[string]bool{}
)

func init() {
	for _, res := range resolutionList {
		ResolutionMap[res] = true
	}
}

// Resolutions removes video resolutions from the the string
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
