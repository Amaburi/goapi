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
	if err := models.DB.Preload("PlayList.Songs").Find(&albums).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch albums: " + err.Error()})
		return
	}

	var result []gin.H
	for _, album := range albums {
		result = append(result, gin.H{
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
			"playlist": gin.H{
				"id":     album.PlayList.ID,
				"name":   album.PlayList.Name,
				"artist": album.PlayList.Artist,
				"likes":  album.PlayList.Likes,
				"saved":  album.PlayList.Saved,
				"songs":  album.PlayList.Songs,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"albums": result})
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

	var playlist models.PlayList
	if album.PlayListID != 0 {
		if err := models.DB.Preload("Songs").First(&playlist, album.PlayListID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch playlist: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"album": gin.H{
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
		},
		"playlist": playlist,
	})
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

func FilterGenres(c *gin.Context) {
	var albums []models.Album

	genre := c.Query("genre")
	if genre == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Please provide a genre to filter by."})
		return
	}

	result := models.DB.Preload("PlayList.Songs").Where("genre = ?", genre).Find(&albums)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch albums: " + result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No albums found for the specified genre."})
		return
	}

	var resultAlbums []gin.H
	for _, album := range albums {
		resultAlbums = append(resultAlbums, gin.H{
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
			"playlist": gin.H{
				"id":     album.PlayList.ID,
				"name":   album.PlayList.Name,
				"artist": album.PlayList.Artist,
				"likes":  album.PlayList.Likes,
				"saved":  album.PlayList.Saved,
				"songs":  album.PlayList.Songs,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"albums": resultAlbums})
}
