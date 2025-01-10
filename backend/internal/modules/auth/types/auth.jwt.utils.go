package types

import (
	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	TOKEN_ISSUER                     = "city-zen"
	ACCESS_TOKEN_HEADER              = "Authorization"
	EMPLOYEE_ACCESS_TOKEN_COOKIE_KEY = "EMPLOYEE_ACCESS_TOKEN"
	USER_ACCESS_TOKEN_COOKIE_KEY     = "USER_ACCESS_TOKEN"
)

type EmployeeClaimsData struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Role         string `json:"role"`
	DepartmentID uint   `json:"departmentId"`
	IsVerified   bool   `json:"isVerified"`
}

type UserClaimsData struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Role       string `json:"role"`
	IsVerified bool   `json:"isVerified"`
}

type EmployeeClaims struct {
	jwt.RegisteredClaims
	EmployeeClaimsData
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserClaimsData
}

func GenerateCustomerJWT(input *UserClaimsData) (string, error) {
	cfg := config.GetConfig()

	claims := UserClaims{
		UserClaimsData: *input,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.AccessTokenExpireMinutes)),
			Issuer:    TOKEN_ISSUER,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenWithClaims.SignedString([]byte(cfg.JWT.UserSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateEmployeeJWT(input *EmployeeClaimsData) (string, error) {
	cfg := config.GetConfig()

	claims := EmployeeClaims{
		EmployeeClaimsData: *input,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.AccessTokenExpireMinutes)),
			Issuer:    TOKEN_ISSUER,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenWithClaims.SignedString([]byte(cfg.JWT.EmployeeSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
