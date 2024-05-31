package repositories

import (
	"btpn-final/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *userRepository) DeleteUser(userID string) error {
	return ur.db.Delete(&models.User{}, userID).Error
}
