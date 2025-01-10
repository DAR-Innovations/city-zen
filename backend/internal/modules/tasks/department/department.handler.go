package department

import (
	"strconv"

	"github.com/DAR-Innovations/city-zen/internal/middleware/contexts"
	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/department/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type DepartmentHandler struct {
	service   DepartmentService
	validator *validator.Validate
}

func NewDepartmentHandler(service DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *DepartmentHandler) GetTasks(c fiber.Ctx) error {
	claims, err := contexts.GetEmployeeClaimsFromCtx(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	filter := types.TaskFilterDTO{
		Status:       c.Query("status"),
		Urgency:      c.Query("urgency"),
		Complexity:   c.Query("complexity"),
		DepartmentID: claims.DepartmentID,
	}

	page := 1   // Default page
	limit := 10 // Default limit

	if pageStr := c.Query("page"); pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil || parsedPage < 1 {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid page parameter")
		}
		page = parsedPage
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < 1 {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid limit parameter")
		}
		limit = parsedLimit
	}

	tasks, err := h.service.GetTasks(filter, page, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(tasks)
}

func (h *DepartmentHandler) AssignTask(c fiber.Ctx) error {
	taskIDstr := c.Params("taskId")
	taskID64, err := strconv.ParseUint(taskIDstr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid task ID")
	}

	var payload types.AssignTaskDTO
	if err := c.Bind().JSON(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid payload")
	}
	if err := h.validator.Struct(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.service.AssignTask(uint(taskID64), payload.AssigneeID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func (h *DepartmentHandler) ReportTask(c fiber.Ctx) error {
	taskIDstr := c.Params("taskId")
	taskID64, err := strconv.ParseUint(taskIDstr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid task ID")
	}

	var payload types.ReportTaskDTO
	if err := c.Bind().JSON(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid payload")
	}
	if err := h.validator.Struct(payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.service.ReportTask(uint(taskID64), &payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
