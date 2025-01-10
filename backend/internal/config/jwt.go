package config

import "time"

type JWTConfig struct {
	UserSecretKey            string        `mapstructure:"USER_SECRET_KEY"`
	EmployeeSecretKey        string        `mapstructure:"EMPLOYEE_SECRET_KEY"`
	AccessTokenExpireMinutes time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRE_MINUTES"`
}
