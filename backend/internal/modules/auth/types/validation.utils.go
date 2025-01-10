package types

import (
	"errors"
	"fmt"
	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func extractAndValidateEmployeeToken(c fiber.Ctx) (*EmployeeClaims, error) {
	tokenString, err := ExtractToken(c, ACCESS_TOKEN_HEADER, EMPLOYEE_ACCESS_TOKEN_COOKIE_KEY)
	if err != nil {
		return nil, err
	}

	claims := &EmployeeClaims{}

	if err := ValidateEmployeeJWT(tokenString, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

func extractAndValidateUserToken(c fiber.Ctx) (*UserClaims, error) {
	tokenString, err := ExtractToken(c, ACCESS_TOKEN_HEADER, USER_ACCESS_TOKEN_COOKIE_KEY)
	if err != nil {
		return nil, err
	}

	claims := &UserClaims{}

	if err := ValidateUserJWT(tokenString, claims); err != nil {
		return nil, err
	}

	return claims, nil
}

// ExtractToken tries to get token from HTTP header or browser cookie. Header is prioritized.
func ExtractToken(c fiber.Ctx, headerKey, cookieKey string) (string, error) {
	// Get from header first
	authHeader := c.Get(headerKey)
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	// Try cookie if header not found
	cookie := c.Cookies(cookieKey)
	if cookie == "" {
		return "", fmt.Errorf("token not found in header or cookie")
	}

	return cookie, nil
}

func ValidateUserJWT(tokenString string, claims *UserClaims) error {
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.UserSecretKey), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	err = validateJWT(token)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	return nil
}

func ValidateEmployeeJWT(tokenString string, claims *EmployeeClaims) error {
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.EmployeeSecretKey), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	err = validateJWT(token)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	return nil
}

func validateJWT(token *jwt.Token) error {
	if !token.Valid {
		return errors.New("token is invalid")
	}

	return nil
}

func ExtractEmployeeAccessTokenAndValidate(c fiber.Ctx) (*EmployeeClaims, error) {
	return extractAndValidateEmployeeToken(c)
}

func ExtractUserAccessTokenAndValidate(c fiber.Ctx) (*UserClaims, error) {
	return extractAndValidateUserToken(c)
}
