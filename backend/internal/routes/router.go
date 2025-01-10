package routes

import (
	"github.com/DAR-Innovations/city-zen/internal/middleware"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth"
	"github.com/DAR-Innovations/city-zen/internal/modules/images"
	"github.com/DAR-Innovations/city-zen/internal/modules/issues"
	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/department"
	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/volunteer"
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

func RegisterTaskRouter(router fiber.Router, departmentHandler department.DepartmentHandler, volunteerHandler volunteer.VolunteerHandler) {
	departmentRouter := router.Group("/department/tasks")
	{
		departmentRouter.Get("", departmentHandler.GetTasks)
		departmentRouter.Post("/:taskId/report", departmentHandler.ReportTask)
		departmentRouter.Post("/:taskId/assign", departmentHandler.AssignTask)
	}

	volunteerRouter := router.Group("/volunteer")
	{
		volunteerRouter.Get("", volunteerHandler.GetTasks)
		volunteerRouter.Post("/:taskId/report", volunteerHandler.ReportTask)
	}
}
