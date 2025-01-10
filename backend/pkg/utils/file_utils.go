package utils

import (
	"mime"
	"path/filepath"
)

func IsValidImage(filename string) bool {
	mimeType := mime.TypeByExtension(filepath.Ext(filename))
	return mimeType == "image/jpeg" || mimeType == "image/png"
}
