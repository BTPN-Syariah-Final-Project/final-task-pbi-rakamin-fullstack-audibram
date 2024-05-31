package usecases

import (
	"btpn-final/dto"
	"btpn-final/models"
	"btpn-final/repositories"
)

type PhotoUsecase interface {
	AddPhoto(req dto.PhotoUploadRequest, userID uint) (*models.Photo, error)
	GetPhotoByID(photoID string) (*models.Photo, error)
	DeletePhoto(photoID string) error
}

type photoUsecase struct {
	photoRepository repositories.PhotoRepository
}

func NewPhotoUsecase(pr repositories.PhotoRepository) PhotoUsecase {
	return &photoUsecase{pr}
}

func (pu *photoUsecase) AddPhoto(req dto.PhotoUploadRequest, userID uint) (*models.Photo, error) {
	photo := models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   userID,
	}

	if err := pu.photoRepository.CreatePhoto(&photo); err != nil {
		return nil, err
	}

	return &photo, nil
}

func (pu *photoUsecase) GetPhotoByID(photoID string) (*models.Photo, error) {
	return pu.photoRepository.GetPhotoByID(photoID)
}

func (pu *photoUsecase) DeletePhoto(photoID string) error {
	return pu.photoRepository.DeletePhoto(photoID)
}
