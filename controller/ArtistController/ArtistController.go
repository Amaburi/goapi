package ArtistController

import (
	"net/http"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var artist []models.Artist
	models.DB.Find(&artist)
	c.JSON(http.StatusOK, gin.H{"Artists": artist})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist
	done := make(chan bool) // Create a channel for synchronization

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()

		if err := models.DB.First(&artist, id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Artist not found"})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			return
		}
	}()

	<-done // Wait for the goroutine to complete
	c.JSON(http.StatusOK, gin.H{"Artist": artist})
}

func Create(c *gin.Context) {
	var artist models.Artist
	done := make(chan bool)
	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		models.DB.Create(&artist)

	}()
	<-done
	c.JSON(http.StatusOK, gin.H{"Artists": artist})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist
	done := make(chan bool)

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		if models.DB.Model(&artist).Where("id = ?", id).Updates(&artist).RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update the Data please try again"})
			return
		}
	}()
	<-done
	c.JSON(http.StatusOK, gin.H{"message": "Data successfully Updated"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	// Perform a DELETE operation based on the company's ID
	if err := models.DB.Where("id = ?", id).Delete(&models.Artist{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artist data deleted successfully"})
}
