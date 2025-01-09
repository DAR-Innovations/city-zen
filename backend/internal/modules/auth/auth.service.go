package auth

import (
	"github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"gorm.io/gorm"
)

type AuthenticationService interface {
	EmployeeSignIn(email, password string) (*types.SignInResponseDTO, error)
	UserSignUp(input *types.UserSignUpRequestDTO) (uint, error)
	UserSignIn(email, password string) (*types.SignInResponseDTO, error)
}

type authenticationService struct {
	db *gorm.DB
}

func NewAuthenticationService(
	db *gorm.DB,
) AuthenticationService {
	return &authenticationService{
		db: db,
	}
}

func (s *authenticationService) EmployeeSignIn(email, password string) (*types.SignInResponseDTO, error) {
	return nil, nil
}
func (s *authenticationService) UserSignUp(input *types.UserSignUpRequestDTO) (uint, error) {
	return 0, nil
}
func (s *authenticationService) UserSignIn(email, password string) (*types.SignInResponseDTO, error) {
	return nil, nil
}
