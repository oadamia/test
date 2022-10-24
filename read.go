package test

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Read reads json data from file
func Read(file string) string {
	content, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer content.Close()
	byteValue, err := io.ReadAll(content)
	if err != nil {
		return ""
	}
	return RemoveEnterAndTabs((string(byteValue)))
}

// RemoveEnterAndTabs removes enters and tabs
func RemoveEnterAndTabs(s string) string {
	return fmt.Sprintf("%s\n", strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), "\t", ""), " ", ""))
}
