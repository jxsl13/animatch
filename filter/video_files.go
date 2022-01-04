package filter

import (
	"path/filepath"
	"strings"
)

var (
	// https://en.wikipedia.org/wiki/Video_file_format
	videoFileExtensionList = []string{
		".3g2",
		".3gp",
		".amv",
		".asf",
		".avi",
		".drc",
		".flv",
		".flv",
		".gif",
		".gifv",
		".m2v",
		".m4p",
		".m4v",
		".m4v",
		".mkv",
		".mng",
		".mov",
		".mp2",
		".mp4",
		".mpe",
		".mpeg",
		".mpeg",
		".mpg",
		".mpg",
		".mpv",
		".mxf",
		".nsv",
		".ogg",
		".ogv",
		".qt",
		".rm",
		".rmvb",
		".roq",
		".svi",
		".viv",
		".vob",
		".webm",
		".wmv",
		".yuv",
	}

	VideoFileExtensionsMap = map[string]bool{}
)

func init() {
	for _, ext := range videoFileExtensionList {
		VideoFileExtensionsMap[ext] = true
	}
}

func VideoFilePaths(vfs []string) []string {
	result := make([]string, 0, len(vfs)/8)
	for _, vf := range vfs {
		lcExt := strings.ToLower(filepath.Ext(vf))
		if VideoFileExtensionsMap[lcExt] {
			result = append(result, vf)
		}
	}
	return result
}
