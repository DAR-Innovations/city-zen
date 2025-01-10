package issues

import (
	"errors"
	"github.com/DAR-Innovations/city-zen/internal/data"
	"github.com/DAR-Innovations/city-zen/internal/modules/issues/types"
	"gorm.io/gorm"
)

type IssuesService struct {
	db *gorm.DB
}

func NewIssuesService(db *gorm.DB) *IssuesService {
	return &IssuesService{db: db}
}

// CreateIssue creates a new issue record in the database
func (s *IssuesService) CreateIssue(authorID uint, dto *types.PostIssueRequestDTO, imagePath string) (uint, error) {
	// Save the issue to the database
	issue := &data.Issue{
		AuthorID:    authorID,
		Name:        dto.Name,
		Description: dto.Description,
		Longitude:   dto.Longitude,
		Latitude:    dto.Latitude,
		ImageURL:    imagePath,
	}

	if err := s.db.Create(&issue).Error; err != nil {
		return 0, err
	}

	return issue.ID, nil
}

// GetIssuesByAuthor retrieves all issues posted by a specific author
func (s *IssuesService) GetIssuesByAuthor(authorID int) ([]types.IssueResponseDTO, error) {
	var issues []data.Issue
	if err := s.db.Where("author_id = ?", authorID).Find(&issues).Error; err != nil {
		return nil, err
	}

	// Convert issues to DTOs
	var response []types.IssueResponseDTO
	for _, issue := range issues {
		response = append(response, types.IssueResponseDTO{
			ID:          issue.ID,
			Name:        issue.Name,
			Description: issue.Description,
			IsCompleted: issue.IsCompleted,
			Longitude:   issue.Longitude,
			Latitude:    issue.Latitude,
			ImageUrl:    issue.ImageURL,
			AuthorID:    issue.AuthorID,
		})
	}
	return response, nil
}

// GetIssueByID retrieves a specific issue by its ID
func (s *IssuesService) GetIssueByID(issueID string) (*types.IssueResponseDTO, error) {
	var issue data.Issue
	if err := s.db.Where("id = ?", issueID).First(&issue).Error; err != nil {
		return nil, errors.New("issue not found")
	}

	// Return the issue as a DTO
	return &types.IssueResponseDTO{
		ID:          issue.ID,
		Name:        issue.Name,
		Description: issue.Description,
		IsCompleted: issue.IsCompleted,
		Longitude:   issue.Longitude,
		Latitude:    issue.Latitude,
		ImageUrl:    issue.ImageURL,
		AuthorID:    issue.AuthorID,
	}, nil
}
