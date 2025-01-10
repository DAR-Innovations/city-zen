package routes

import (
	"github.com/DAR-Innovations/city-zen/internal/middleware"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth"
	"github.com/DAR-Innovations/city-zen/internal/modules/images"
	"github.com/DAR-Innovations/city-zen/internal/modules/issues"
	"github.com/gofiber/fiber/v3"
)

func RegisterAuthRoutes(router fiber.Router, handler auth.AuthenticationHandler) {
	authRouter := router.Group("/auth")
	authRouter.Post("/user/signin", handler.UserSignIn)
	authRouter.Post("/user/signup", handler.UserSignUp)
	authRouter.Get("/user/me", handler.CurrentUser, middleware.UserAuth())
	authRouter.Post("/employee/signin", handler.EmployeeSignIn)
	authRouter.Post("/employee/create", handler.CreateEmployee, middleware.EmployeeAuth())
	authRouter.Get("/employee/me", handler.CurrentEmployee, middleware.EmployeeAuth())
}

func RegisterIssueRoutes(router fiber.Router, handler issues.IssuesHandler) {
	issueRouter := router.Group("/issues")
	issueRouter.Post("/", handler.CreateIssue, middleware.UserAuth())
	issueRouter.Get("/my", handler.GetMyIssues, middleware.UserAuth())
	issueRouter.Get("/:id", handler.GetIssueByID)
}

func RegisterImagesRoutes(router fiber.Router) {
	imagesRouter := router.Group("/images")
	imagesRouter.Post("/upload", images.UploadImage)
}
