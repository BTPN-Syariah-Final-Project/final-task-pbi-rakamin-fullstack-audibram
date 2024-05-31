package usecases

import (
	"btpn-final/helpers"
	"btpn-final/models"
	"btpn-final/repositories"
	"errors"
)

type UserUsecase interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (*models.User, string, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID uint) error
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (u *userUsecase) RegisterUser(user *models.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) LoginUser(email, password string) (*models.User, string, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(userID uint) error {
	return u.userRepository.DeleteUser(userID)
}
