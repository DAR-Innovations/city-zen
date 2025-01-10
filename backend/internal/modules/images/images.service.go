package images

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	OriginalFilename string    `json:"original_filename"`
	ContentType      string    `json:"content_type"`
	FileSize         int64     `json:"file_size"`
	UploadedAt       time.Time `json:"uploaded_at"`
}

func SaveImageWithMetadata(fileName string, content []byte, uploadFolder string) (string, string, error) {
	folderName := uuid.New().String()
	folderPath := filepath.Join(uploadFolder, folderName)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("failed to create folder: %w", err)
	}

	imageFileName := uuid.New().String() + filepath.Ext(fileName)
	imagePath := filepath.Join(folderPath, imageFileName)
	if err := os.WriteFile(imagePath, content, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("failed to save image: %w", err)
	}

	metadata := Metadata{
		OriginalFilename: fileName,
		ContentType:      "image",
		FileSize:         int64(len(content)),
		UploadedAt:       time.Now(),
	}

	metadataPath := filepath.Join(folderPath, "metadata.json")
	metadataFile, err := os.Create(metadataPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to create metadata file: %w", err)
	}
	defer metadataFile.Close()

	if err := json.NewEncoder(metadataFile).Encode(metadata); err != nil {
		return "", "", fmt.Errorf("failed to write metadata: %w", err)
	}

	return imagePath, metadataPath, nil
}
