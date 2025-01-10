package issues

import (
	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/DAR-Innovations/city-zen/internal/middleware/contexts"
	"github.com/DAR-Innovations/city-zen/internal/modules/images"
	"github.com/DAR-Innovations/city-zen/internal/modules/issues/types"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"io"
)

type IssuesHandler struct {
	service IssuesService
}

func NewIssuesHandler(service IssuesService) *IssuesHandler {
	return &IssuesHandler{service: service}
}

// CreateIssue creates a new issue
func (h *IssuesHandler) CreateIssue(c fiber.Ctx) error {
	var dto types.PostIssueRequestDTO

	// Bind the request form data to DTO
	if err := c.Bind().Form(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Extract the image from the form data
	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Image is required",
		})
	}

	file, err := image.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to open file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to read file")
	}

	fileName := uuid.New().String()

	uploadFolder := config.GetConfig().UploadFolder
	if uploadFolder == "" {
		uploadFolder = "./uploads"
	}

	// Save the image and extract metadata
	filePath, _, err := images.SaveImageWithMetadata(fileName, content, uploadFolder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save image with metadata",
		})
	}

	claims, err := contexts.GetUserClaimsFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user claims",
		})
	}

	// Call service to create issue
	issueResponse, err := h.service.CreateIssue(claims.UserClaimsData.ID, &dto, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(issueResponse)
}

// GetMyIssues returns all issues for the current user
func (h *IssuesHandler) GetMyIssues(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)

	issues, err := h.service.GetIssuesByAuthor(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(issues)
}

// GetIssueByID returns a specific issue by ID
func (h *IssuesHandler) GetIssueByID(c fiber.Ctx) error {
	issueID := c.Params("id")
	issue, err := h.service.GetIssueByID(issueID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Issue not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(issue)
}
