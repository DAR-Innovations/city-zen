package utils

import (
	"mime"
	"path/filepath"
)

// IsValidImage checks if the file is a valid image type
func IsValidImage(filename string) bool {
	mimeType := mime.TypeByExtension(filepath.Ext(filename))
	return mimeType == "image/jpeg" || mimeType == "image/png"
}
