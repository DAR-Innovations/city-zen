package types

type UserSignUpRequestDTO struct {
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required,min=4"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type SignInRequestDTO struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=4"`
}

type CreateEmployeeRequestDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Password  string `json:"password" validate:"required,min=4"`
	Role      string `json:"role" validate:"required"`
}

type UserResponseDTO struct {
	ID         int    `json:"id"`
	Phone      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Role       string `json:"role"`
	IsVerified bool   `json:"isVerified"`
}

type EmployeeResponseDTO struct {
	ID           int    `json:"id"`
	Phone        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DepartmentID string `json:"departmentId"`
	Role         string `json:"role"`
	IsVerified   bool   `json:"isVerified"`
}

type SignInResponseDTO struct {
	AccessToken string `json:"accessToken"`
}
