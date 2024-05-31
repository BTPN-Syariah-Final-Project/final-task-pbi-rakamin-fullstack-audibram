package controllers

import (
	"btpn-final/models"
	"btpn-final/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PhotoController struct {
	photoUsecase usecases.PhotoUsecase
}

func NewPhotoController(photoUsecase usecases.PhotoUsecase) *PhotoController {
	return &PhotoController{photoUsecase: photoUsecase}
}

func (pc *PhotoController) AddPhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.photoUsecase.AddPhoto(&photo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo added successfully"})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	userID, _ := c.Get("userID")
	photoIDStr := c.Param("photoID")

	photoID, err := strconv.ParseUint(photoIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid photoID"})
		return
	}

	if err := pc.photoUsecase.DeletePhoto(uint(photoID), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo deleted successfully"})
}

func (p *PhotoController) GetPhoto(c *gin.Context) {
	photoID := c.Param("photoID")

	photo, err := p.photoUsecase.GetPhotoByID(photoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}
