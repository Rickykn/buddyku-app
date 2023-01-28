package services

import (
	"github.com/Rickykn/buddyku-app.git/dtos"
	help "github.com/Rickykn/buddyku-app.git/helpers"
	"github.com/Rickykn/buddyku-app.git/models"
	r "github.com/Rickykn/buddyku-app.git/repositories"
)

type AdminService interface {
	// SetPoint()
	RegisterAdmin(registerAdminInput *dtos.AdminRegisterDTO) (*models.Admin, *help.JsonResponse, error)
	// Login(loginInput *dtos.LoginUserDTO) (*dtos.TokenDTO, *help.JsonResponse, error)
}

type adminService struct {
	adminRepository r.AdminRepository
}

type ASConfig struct {
	AdminRespository r.AdminRepository
}

func NewAdminService(c *ASConfig) AdminService {
	return &adminService{
		adminRepository: c.AdminRespository,
	}
}

func (a *adminService) RegisterAdmin(registerAdminInput *dtos.AdminRegisterDTO) (*models.Admin, *help.JsonResponse, error) {
	_, row, err := a.adminRepository.FindOneAdmin(registerAdminInput.Name)

	if row == 1 {
		return nil, help.HandlerError(400, "Email has been Taken", nil), err
	}

	hasingPassword, _ := help.HashPassword(registerAdminInput.Password)

	newAdmin, err := a.adminRepository.CreateAdmin(registerAdminInput.Name, hasingPassword)

	if err != nil {
		return nil, help.HandlerError(500, "Server Error", nil), err
	}

	return newAdmin, help.HandlerSuccess(201, "Success register admin", newAdmin), nil
}
