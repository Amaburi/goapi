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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body: " + err.Error()})
		return
	}

	if artist.Name == "" || artist.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Name and Email are required"})
		return
	}

	if artist.CompanyID != 0 {
		var company models.Company
		if err := models.DB.First(&company, artist.CompanyID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid company ID"})
			return
		}
	}

	if len(artist.Albums) > 0 {
		for i, album := range artist.Albums {
			if album.ID == 0 {
				if err := models.DB.Create(&album).Error; err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create album: " + err.Error()})
					return
				}
				artist.Albums[i] = album
			} else {
				var existingAlbum models.Album
				if err := models.DB.First(&existingAlbum, album.ID).Error; err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid album ID: " + err.Error()})
					return
				}
			}
		}
	}

	if err := models.DB.Create(&artist).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create artist: " + err.Error()})
		return
	}

	if len(artist.Albums) > 0 {
		if err := models.DB.Model(&artist).Association("Albums").Append(artist.Albums); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to associate albums: " + err.Error()})
			return
		}
	}

	if err := models.DB.Preload("Labels").First(&artist, artist.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch artist with labels: " + err.Error()})
		return
	}

	albumsResponse := make([]gin.H, len(artist.Albums))
	for i, album := range artist.Albums {
		albumsResponse[i] = gin.H{
			"id":          album.ID,
			"title":       album.Title,
			"artist":      album.Artist,
			"price":       album.Price,
			"playlist_id": album.PlayListID,
			"description": album.Description,
			"awards":      album.Awards,
			"genre":       album.Genre,
			"releasedate": album.Relasedate,
			"rating":      album.Rating,
			"link":        album.Link,
			"cover_art":   album.CoverArt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"artist": gin.H{
			"id":                   artist.ID,
			"name":                 artist.Name,
			"age":                  artist.Age,
			"address":              artist.Address,
			"phone":                artist.PhoneNumber,
			"social_media_account": artist.SocialMediaAccount,
			"achievement":          artist.Achievement,
			"biography":            artist.Biography,
			"nationality":          artist.Nationality,
			"website":              artist.Website,
			"email":                artist.Email,
			"labels":               artist.Labels,
			"company_id":           artist.CompanyID,
			"albums":               albumsResponse,
			"profile_picture":      artist.ProfilePicture,
		},
	})
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
