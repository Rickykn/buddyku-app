package repositories

import (
	"github.com/Rickykn/buddyku-app.git/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(email, name, password string) (*models.User, error)

	FindOneUser(email string) (*models.User, int, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (u *userRepository) CreateUser(email, name, password string) (*models.User, error) {
	newUser := &models.User{
		Name:         name,
		Email:        email,
		Password:     password,
		Point_reward: 0,
	}

	result := u.db.Create(&newUser)

	return newUser, result.Error
}
func (u *userRepository) FindOneUser(email string) (*models.User, int, error) {
	var user *models.User

	result := u.db.Where("email = ?", email).First(&user)

	return user, int(result.RowsAffected), result.Error
}
