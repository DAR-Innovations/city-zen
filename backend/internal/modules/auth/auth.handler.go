package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type AuthenticationHandler struct {
	service   AuthenticationService
	validator *validator.Validate
}

func NewAuthenticationHandler(service AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *AuthenticationHandler) UserSignIn(c fiber.Ctx) error {
	return nil
}
