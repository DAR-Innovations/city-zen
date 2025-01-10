package auth

import (
	"errors"
	"fmt"
	"github.com/DAR-Innovations/city-zen/internal/data"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth/types"
	"github.com/DAR-Innovations/city-zen/pkg/utils"
	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

type AuthenticationService interface {
	EmployeeSignIn(dto *types.SignInRequestDTO) (*types.SignInResponseDTO, error)
	CreateEmployee(departmentID uint, dto *types.CreateEmployeeRequestDTO) (uint, error)
	UserSignUp(input *types.UserSignUpRequestDTO) (uint, error)
	UserSignIn(dto *types.SignInRequestDTO) (*types.SignInResponseDTO, error)
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

func (s *authenticationService) EmployeeSignIn(dto *types.SignInRequestDTO) (*types.SignInResponseDTO, error) {
	var employee data.Employee

	if err := s.db.Where("phone = ?", dto.Phone).First(&employee).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid phone or password")
		}
		return nil, err
	}

	if err := utils.ComparePassword(employee.Password, dto.Password); err != nil {
		return nil, errors.New("invalid phone or password")
	}

	token, err := types.GenerateEmployeeJWT(&types.EmployeeClaimsData{
		ID:           employee.ID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Role:         employee.Role,
		DepartmentID: *employee.DepartmentID,
		IsVerified:   employee.IsVerified,
	})
	if err != nil {
		return nil, err
	}

	return &types.SignInResponseDTO{
		AccessToken: token,
	}, nil
}

func (s *authenticationService) CreateEmployee(departmentID uint, dto *types.CreateEmployeeRequestDTO) (uint, error) {
	employee := &data.Employee{
		DepartmentID: &departmentID,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Role:         dto.Role,
		Phone:        dto.Phone,
		Password:     dto.Password,
		IsVerified:   false,
	}

	if err := s.db.Create(employee).Error; err != nil {
		return 0, err
	}

	return employee.ID, nil
}

func (s *authenticationService) UserSignUp(dto *types.UserSignUpRequestDTO) (uint, error) {
	if err := utils.IsValidPassword(dto.Password); err != nil {
		return 0, err
	}

	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return 0, err
	}

	user := data.User{
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		Phone:      dto.Phone,
		Role:       "USER",
		Password:   hashedPassword,
		IsVerified: false,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *authenticationService) UserSignIn(dto *types.SignInRequestDTO) (*types.SignInResponseDTO, error) {
	var user data.User

	if err := s.db.Where("phone = ?", dto.Phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error()
			return nil, errors.New(fmt.Sprintf("invalid phone or password phone: %s %s", dto.Phone, dto.Password))
		}
		return nil, err
	}

	if err := utils.ComparePassword(user.Password, dto.Password); err != nil {
		return nil, errors.New("invalid phone or password")
	}

	token, err := types.GenerateCustomerJWT(&types.UserClaimsData{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Role:       user.Role,
		IsVerified: user.IsVerified,
	})
	if err != nil {
		return nil, err
	}

	return &types.SignInResponseDTO{
		AccessToken: token,
	}, nil
}
