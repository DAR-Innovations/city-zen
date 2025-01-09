package ai

import "github.com/DAR-Innovations/city-zen/internal/config"

type ValidationResult struct {
	IsValid bool   `json:"is_valid"`
	Reason  string `json:"reason,omitempty"`
}

var cfg = config.GetConfig()
