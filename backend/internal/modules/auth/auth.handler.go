package auth

import (
	"github.com/DAR-Innovations/city-zen/internal/middleware/contexts"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
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
	var dto types.SignInRequestDTO

	if err := c.Bind().JSON(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.service.UserSignIn(&dto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := &fiber.Cookie{
		Name:     types.USER_ACCESS_TOKEN_COOKIE_KEY,
		Value:    response.AccessToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	}

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthenticationHandler) EmployeeSignIn(c fiber.Ctx) error {
	var dto types.SignInRequestDTO

	if err := c.Bind().JSON(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		logrus.Error("validation error")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response, err := h.service.EmployeeSignIn(&dto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := &fiber.Cookie{
		Name:     types.EMPLOYEE_ACCESS_TOKEN_COOKIE_KEY,
		Value:    response.AccessToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	}

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthenticationHandler) UserSignUp(c fiber.Ctx) error {
	var dto types.UserSignUpRequestDTO

	if err := c.Bind().JSON(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := h.service.UserSignUp(&dto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (h *AuthenticationHandler) CreateEmployee(c fiber.Ctx) error {
	var dto types.CreateEmployeeRequestDTO

	if err := c.Bind().JSON(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims, err := contexts.GetEmployeeClaimsFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if claims.Role != "ADMIN" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You do not have permission to create employees",
		})
	}

	id, err := h.service.CreateEmployee(claims.DepartmentID, &dto)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (h *AuthenticationHandler) CurrentUser(c fiber.Ctx) error {

	claims, err := contexts.GetUserClaimsFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(claims.UserClaimsData)
}

func (h *AuthenticationHandler) CurrentEmployee(c fiber.Ctx) error {
	claims, err := contexts.GetEmployeeClaimsFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(claims.EmployeeClaimsData)
}
