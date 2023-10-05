package CompanyController

import (
	"net/http"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var company []models.Company
	models.DB.Find(&company)
	c.JSON(http.StatusOK, gin.H{"Company": company})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	done := make(chan bool) // Create a channel for synchronization

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()

		if err := models.DB.First(&company, id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Company not found"})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			return
		}
	}()

	<-done // Wait for the goroutine to complete
	c.JSON(http.StatusOK, gin.H{"Company": company})
}

func Create(c *gin.Context) {
	var company models.Company
	done := make(chan bool)
	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		models.DB.Create(&company)

	}()
	<-done
	c.JSON(http.StatusOK, gin.H{"Company": company})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	done := make(chan bool)

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	go func() {
		defer func() {
			done <- true // Signal that the goroutine has completed
		}()
		if models.DB.Model(&company).Where("id = ?", id).Updates(&company).RowsAffected == 0 {
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
	if err := models.DB.Where("id = ?", id).Delete(&models.Company{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company data deleted successfully"})
}
