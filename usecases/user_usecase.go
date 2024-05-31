package usecases

import (
	"btpn-final/dto"
	"btpn-final/helpers"
	"btpn-final/models"
	"btpn-final/repositories"
	"errors"
)

type UserUsecase interface {
	RegisterUser(req dto.UserRegisterRequest) (*models.User, error)
	LoginUser(req dto.UserLoginRequest) (string, error)
	GetUserByID(userID string) (*models.User, error)
	UpdateUser(userID string, req dto.UserUpdateRequest) (*models.User, error)
	DeleteUser(userID string) error
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(ur repositories.UserRepository) UserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) RegisterUser(req dto.UserRegisterRequest) (*models.User, error) {
	hashedPassword, _ := helpers.HashPassword(req.Password)

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := uu.userRepository.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (uu *userUsecase) LoginUser(req dto.UserLoginRequest) (string, error) {
	user, err := uu.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if !helpers.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uu *userUsecase) GetUserByID(userID string) (*models.User, error) {
	return uu.userRepository.GetUserByID(userID)
}

func (uu *userUsecase) UpdateUser(userID string, req dto.UserUpdateRequest) (*models.User, error) {
	user, err := uu.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, _ := helpers.HashPassword(req.Password)

		user.Password = hashedPassword
	}

	if err := uu.userRepository.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) DeleteUser(userID string) error {
	return uu.userRepository.DeleteUser(userID)
}
