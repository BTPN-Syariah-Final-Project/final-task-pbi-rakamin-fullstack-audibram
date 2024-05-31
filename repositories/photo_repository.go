package repositories

import (
	"btpn-final/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) error
	GetPhotoByID(photoID string) (*models.Photo, error)
	DeletePhoto(photoID uint) error
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db}
}

func (p *photoRepository) CreatePhoto(photo *models.Photo) error {
	return p.db.Create(photo).Error
}

func (p *photoRepository) GetPhotoByID(photoID string) (*models.Photo, error) {
	var photo models.Photo
	if err := p.db.First(&photo, photoID).Error; err != nil {
		return nil, err
	}
	return &photo, nil
}

func (p *photoRepository) DeletePhoto(photoID uint) error {
	return p.db.Delete(&models.Photo{}, photoID).Error
}
