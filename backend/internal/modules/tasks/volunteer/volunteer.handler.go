package volunteer

import (
	"strconv"

	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/volunteer/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type VolunteerHandler struct {
	service   VolunteerService
	validator *validator.Validate
}

func NewVolunteerHandler(service VolunteerService) *VolunteerHandler {
	return &VolunteerHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *VolunteerHandler) GetTasks(c fiber.Ctx) error {
	filter := types.TaskFilterDTO{
		Status:     c.Query("status"),
		Urgency:    c.Query("urgency"),
		Complexity: c.Query("complexity"),
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	tasks, err := h.service.GetTasks(filter, page, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(tasks)
}

func (h *VolunteerHandler) ReportTask(c fiber.Ctx) error {
	taskIDStr := c.Params("taskId")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid task ID")
	}

	var payload types.ReportTaskDTO
	if err := c.Bind().JSON(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid payload")
	}
	if err := h.validator.Struct(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.service.ReportTask(uint(taskID), &payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
