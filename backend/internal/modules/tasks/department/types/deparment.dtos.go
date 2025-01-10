package types

type DepartmentTaskDTO struct {
	ID             uint         `json:"id"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	DepartmentName string       `json:"departmentName"`
	Type           string       `json:"type"`
	ImageURL       string       `json:"imageUrl"`
	Location       LocationDTO  `json:"location"`
	Status         string       `json:"status"`
	Assignee       *AssigneeDTO `json:"assignee,omitempty"`
	Urgency        string       `json:"urgency"`
	Complexity     string       `json:"complexity"`
}

type LocationDTO struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type AssigneeDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type AssignTaskDTO struct {
	AssigneeID uint `json:"assigneeId" validate:"required"`
}

type ReportTaskDTO struct {
	ImageURL string `json:"imageUrl" validate:"required"`
	Comment  string `json:"comment" validate:"required"`
}

type TaskFilterDTO struct {
	Status       string `json:"status,omitempty"`
	Urgency      string `json:"urgency,omitempty"`
	Complexity   string `json:"complexity,omitempty"`
	DepartmentID uint   `json:"departmentId,omitempty"`
}
