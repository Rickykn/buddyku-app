package repositories

import (
	"github.com/Rickykn/buddyku-app.git/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(name, password string) (*models.Admin, error)
	FindOneAdmin(email string) (*models.Admin, int, error)
}

type adminRepository struct {
	db *gorm.DB
}

type ARConfig struct {
	DB *gorm.DB
}

func NewAdminRepository(c *ARConfig) AdminRepository {
	return &adminRepository{
		db: c.DB,
	}
}

func (u *adminRepository) CreateAdmin(name, password string) (*models.Admin, error) {
	newAdmin := &models.Admin{
		Name:     name,
		Password: password,
		Role:     "admin",
	}

	result := u.db.Create(&newAdmin)

	return newAdmin, result.Error
}
func (u *adminRepository) FindOneAdmin(email string) (*models.Admin, int, error) {
	var admin *models.Admin

	result := u.db.Where("email = ?", email).First(&admin)

	return admin, int(result.RowsAffected), result.Error
}
