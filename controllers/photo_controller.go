package controllers

import (
	"btpn-final/dto"
	"btpn-final/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PhotoController struct {
	photoUsecase usecases.PhotoUsecase
}

func NewPhotoController(pu usecases.PhotoUsecase) *PhotoController {
	return &PhotoController{pu}
}

func (pc *PhotoController) AddPhoto(c *gin.Context) {
	var req dto.PhotoUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	photo, err := pc.photoUsecase.AddPhoto(req, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	})
}

func (pc *PhotoController) GetPhoto(c *gin.Context) {
	photoID := c.Param("photoID")

	photo, err := pc.photoUsecase.GetPhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoID")

	err := pc.photoUsecase.DeletePhoto(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
