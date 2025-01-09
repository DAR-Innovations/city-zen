package types

type UserSignUpRequestDTO struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type SignInRequestDTO struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Phone      string `json:"email"`
	Name       string `json:"name"`
	IsVerified bool   `json:"isVerified"`
}

type SignInResponseDTO struct {
	AccessToken string `json:"accessToken"`
}
