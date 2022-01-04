package common

import (
	"os"
	"path"
)

func Readdir(dirPath string) ([]string, error) {
	result := make([]string, 0, 2)
	de, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, d := range de {
		filePath := path.Join(dirPath, d.Name())
		if d.IsDir() {
			filesPaths, _ := Readdir(filePath)
			result = append(result, filesPaths...)
		} else {
			result = append(result, filePath)
		}
	}
	return result, nil
}
