package PlaylistController

import (
	"net/http"
	"strconv"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var playlists []models.PlayList
	if err := models.DB.Find(&playlists).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"playlists": playlists})
}

func Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	var playlist models.PlayList
	if err := models.DB.First(&playlist, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Playlist not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"playlist": playlist})
}

func Create(c *gin.Context) {
	var playlist models.PlayList
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := models.DB.Create(&playlist).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"playlist": playlist})
}

func Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	var playlist models.PlayList
	if err := c.ShouldBindJSON(&playlist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Model(&playlist).Where("id = ?", id).Updates(&playlist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Playlist not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist has been successfully updated"})
}

func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", id).Delete(&models.PlayList{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist has been deleted successfully"})
}
