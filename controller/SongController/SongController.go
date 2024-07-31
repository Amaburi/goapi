package SongController

import (
	"net/http"
	"strconv"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var songs []models.Song
	if err := models.DB.Find(&songs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songss": songs})
}

func Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	var songs models.Song
	if err := models.DB.First(&songs, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "songs not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func Create(c *gin.Context) {
	var songs models.Song
	if err := c.ShouldBindJSON(&songs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := models.DB.Create(&songs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"songs": songs})
}

func Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	var songs models.Song
	if err := c.ShouldBindJSON(&songs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&songs).Where("id = ?", id).Updates(&songs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "song not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song has been successfully updated"})
}

func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", id).Delete(&models.Song{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song has been deleted successfully"})
}
