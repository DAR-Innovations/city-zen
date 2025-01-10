package types

type PostIssueRequestDTO struct {
	Name        string  `form:"name" validate:"required"`
	Description string  `form:"description" validate:"required"`
	Longitude   float64 `form:"longitude" validate:"required"`
	Latitude    float64 `form:"latitude" validate:"required"`
}

type IssueResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsCompleted bool    `json:"isCompleted"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	AuthorID    uint    `json:"authorId"`
	ImageUrl    string  `json:"imageUrl"`
}
