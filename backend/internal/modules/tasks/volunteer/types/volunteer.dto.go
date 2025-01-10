package types

type VolunteerTaskDTO struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ImageURL    string       `json:"imageUrl"`
	Location    LocationDTO  `json:"location"`
	Status      string       `json:"status"`
	Urgency     string       `json:"urgency"`
	Complexity  string       `json:"complexity"`
	Assignee    *AssigneeDTO `json:"assignee,omitempty"`
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

type TaskFilterDTO struct {
	Status     string `json:"status,omitempty"`
	Urgency    string `json:"urgency,omitempty"`
	Complexity string `json:"complexity,omitempty"`
}

type ReportTaskDTO struct {
	ImageURL string `json:"imageUrl" validate:"required"`
	Comment  string `json:"comment" validate:"required"`
}
