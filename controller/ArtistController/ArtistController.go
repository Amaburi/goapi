package ArtistController

import (
	"net/http"
	"strconv"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var artists []models.Artist
	if err := models.DB.Find(&artists).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"artists": artists})
}

func Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var artist models.Artist
	if err := models.DB.First(&artist, uint(id)).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"artist": artist})
}

func Create(c *gin.Context) {
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := models.DB.Create(&artist).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"artist": artist})
}

func Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var artist models.Artist
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Model(&artist).Where("id = ?", uint(id)).Updates(&artist)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Artist data successfully updated"})
}

func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := models.DB.Where("id = ?", uint(id)).Delete(&models.Artist{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artist data deleted successfully"})
}
