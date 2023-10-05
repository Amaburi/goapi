package controller

import (
	"net/http"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"album": albums})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	done := make(chan bool) // Create a channel for synchronization

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()

		if err := models.DB.First(&album, id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			return
		}
	}()

	<-done // Wait for the goroutine to complete
	c.JSON(http.StatusOK, gin.H{"album": album})
}

func Create(c *gin.Context) {
	var album models.Album
	done := make(chan bool)
	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		models.DB.Create(&album)
	}()
	<-done
	c.JSON(http.StatusOK, gin.H{"album": album})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	done := make(chan bool)

	if err := c.ShouldBindJSON(&album); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		if models.DB.Model(&album).Where("id = ?", id).Updates(&album).RowsAffected == 0 {
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
	if err := models.DB.Where("id = ?", id).Delete(&models.Album{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
