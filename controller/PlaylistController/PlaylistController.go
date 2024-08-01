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
	if err := models.DB.Preload("Songs").Find(&playlists).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"playlists": playlists})
}

func Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var playlist models.PlayList
	if err := models.DB.Preload("Songs").First(&playlist, uint(id)).Error; err != nil {
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
	var input struct {
		Name   string `json:"name"`
		Artist string `json:"artist"`
		Likes  uint   `json:"likes"`
		Saved  uint   `json:"saved"`
		Songs  []struct {
			ID uint `json:"id"`
		} `json:"songs"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var songIDs []uint
	for _, s := range input.Songs {
		songIDs = append(songIDs, s.ID)
	}

	var songs []models.Song
	if err := models.DB.Where("id IN ?", songIDs).Find(&songs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch songs: " + err.Error()})
		return
	}

	playlist := models.PlayList{
		Name:   input.Name,
		Artist: input.Artist,
		Likes:  input.Likes,
		Saved:  input.Saved,
		Songs:  songs,
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var input struct {
		Name   string `json:"name"`
		Artist string `json:"artist"`
		Likes  uint   `json:"likes"`
		Saved  uint   `json:"saved"`
		Songs  []struct {
			ID uint `json:"id"`
		} `json:"songs"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var playlist models.PlayList
	if err := models.DB.Preload("Songs").First(&playlist, uint(id)).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Playlist not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	var songIDs []uint
	for _, s := range input.Songs {
		songIDs = append(songIDs, s.ID)
	}

	var songs []models.Song
	if err := models.DB.Where("id IN ?", songIDs).Find(&songs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch songs: " + err.Error()})
		return
	}

	playlist.Name = input.Name
	playlist.Artist = input.Artist
	playlist.Likes = input.Likes
	playlist.Saved = input.Saved
	playlist.Songs = songs

	if err := models.DB.Save(&playlist).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"playlist": playlist})
}

func Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	if err := models.DB.Where("id = ?", id).Delete(&models.PlayList{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Playlist has been deleted successfully"})
}
