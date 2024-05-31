package repositories

import (
	"btpn-final/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) UpdateUser(user *models.User) error {
	return u.db.Save(user).Error
}

func (u *userRepository) DeleteUser(userId uint) error {
	return u.db.Delete(&models.User{}, userId).Error
}
