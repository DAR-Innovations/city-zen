package images

import (
	"fmt"
	"io"

	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/gofiber/fiber/v3"
)

func UploadImage(c fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "File is required")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to open file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to read file")
	}

	uploadFolder := config.GetConfig().UploadFolder
	if uploadFolder == "" {
		uploadFolder = "./uploads"
	}

	imagePath, metadataPath, err := SaveImageWithMetadata(fileHeader.Filename, content, uploadFolder)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to save file: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Image and metadata saved successfully",
		"image_path":    imagePath,
		"metadata_path": metadataPath,
	})
}
