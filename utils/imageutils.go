package utils

import (
	"path/filepath"
	"strings"
)

func HasImageExtension(filename string) bool {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	imageExtensions := []string{"png", "gif", "jpg", "jpeg"}
	for _, imageExtension := range imageExtensions {
		if imageExtension == ext {
			return true
		}
	}
	return false
}