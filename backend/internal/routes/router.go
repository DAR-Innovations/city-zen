package routes

import (
	"github.com/DAR-Innovations/city-zen/internal/modules/auth"
	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router, handler auth.AuthenticationHandler) {
	// User Reports
	authRouter := router.Group("/auth")
	authRouter.Post("/user/signin", handler.UserSignIn)
	/*auth.Post("/user/signup", handler.UserSignUp)
	auth.Post("/employee/signin", handler.EmployeeSignIn)*/
}
