package leetcode

import (
	"strings"
)

func FindDuplicateFileInSystem(paths []string) [][]string {
	return findDuplicate(paths)
}

func findDuplicate(paths []string) [][]string {
	fileMap := make(map[string][]string)

	for _, path := range paths {
		pathAndFiles := strings.Split(path, " ")
		realPath := pathAndFiles[0]
		files := pathAndFiles[1:]

		for _, file := range files {
			fileNameAndContent := strings.Split(file, "(")
			fileName := fileNameAndContent[0]
			content := strings.TrimRight(fileNameAndContent[1], ")")
			fileMap[content] = append(fileMap[content], realPath + "/" + fileName)
		}
	}

	result := make([][]string, 0)

	for _, val := range fileMap {
		if len(val) > 1 {
			result = append(result, val)
		}
	}

	return result
}