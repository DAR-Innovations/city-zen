package middleware

import (
	"github.com/DAR-Innovations/city-zen/internal/middleware/contexts"
	authTypes "github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func EmployeeAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		claims, err := authTypes.ExtractEmployeeAccessTokenAndValidate(c)
		if err != nil {
			logrus.Warn("missing or invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid token",
			})
		}

		contexts.SetEmployeeCtx(c, claims)
		return c.Next()
	}
}

func UserAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		claims, err := authTypes.ExtractUserAccessTokenAndValidate(c)
		if err != nil {
			logrus.Warn("missing or invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid token",
			})
		}

		contexts.SetUserCtx(c, claims)
		return c.Next()
	}
}
