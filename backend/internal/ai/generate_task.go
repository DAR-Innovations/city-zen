package ai

type TaskMetadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Urgency     string `json:"urgency"`
	Complexity  string `json:"complexity"`
	Department  string `json:"department,omitempty"`
}
