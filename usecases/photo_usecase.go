package usecases

import (
	"btpn-final/models"
	"btpn-final/repositories"
	"errors"
	"strconv"
)

type PhotoUsecase interface {
	AddPhoto(photo *models.Photo) error
	DeletePhoto(photoID, userID uint) error
	GetPhotoByID(photoID string) (*models.Photo, error)
}

type photoUsecase struct {
	photoRepository repositories.PhotoRepository
	userRepository  repositories.UserRepository
}

func NewPhotoUsecase(pr repositories.PhotoRepository, ur repositories.UserRepository) PhotoUsecase {
	return &photoUsecase{pr, ur}
}

func (p *photoUsecase) AddPhoto(photo *models.Photo) error {
	return p.photoRepository.CreatePhoto(photo)
}

func (p *photoUsecase) DeletePhoto(photoID, userID uint) error {
	photo, err := p.photoRepository.GetPhotoByID(strconv.Itoa(int(photoID)))
	if err != nil {
		return err
	}

	if photo.UserID != userID {
		return errors.New("wrong user")
	}

	return p.photoRepository.DeletePhoto(photoID)
}

func (p *photoUsecase) GetPhotoByID(photoID string) (*models.Photo, error) {
	return p.photoRepository.GetPhotoByID(photoID)
}
