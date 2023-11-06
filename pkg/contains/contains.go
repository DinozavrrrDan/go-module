package contains

import (
	"os"
	"strings"
)

func Contains(filePath string, substring string) (bool, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	contentString := string(fileContent)
	if strings.Contains(contentString, substring) {
		return true, nil
	}
	return false, nil
}
