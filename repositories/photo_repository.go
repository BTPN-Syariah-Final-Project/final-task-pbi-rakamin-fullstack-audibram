package repositories

import (
	"btpn-final/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) error
	GetPhotoByID(photoID string) (*models.Photo, error)
	DeletePhoto(photoID string) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db}
}

func (pr *photoRepository) CreatePhoto(photo *models.Photo) error {
	return pr.db.Create(photo).Error
}

func (pr *photoRepository) GetPhotoByID(photoID string) (*models.Photo, error) {
	var photo models.Photo
	if err := pr.db.First(&photo, photoID).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}

func (pr *photoRepository) DeletePhoto(photoID string) error {
	return pr.db.Delete(&models.Photo{}, photoID).Error
}
