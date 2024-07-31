package controller

import (
	"net/http"
	"strconv"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var albums []models.Album
	if err := models.DB.Find(&albums).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch albums: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"albums": albums})
}

func Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}
	var album models.Album
	if err := models.DB.First(&album, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch album: " + err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"album": album})
}

func Create(c *gin.Context) {
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body: " + err.Error()})
		return
	}
	if err := models.DB.Create(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create album: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"album": album})
}

func Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}
	var album models.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body: " + err.Error()})
		return
	}

	result := models.DB.Model(&album).Where("id = ?", id).Updates(&album)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update album: " + result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album successfully updated", "album": album})
}

func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}
	if err := models.DB.Where("id = ?", id).Delete(&models.Album{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete album: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
