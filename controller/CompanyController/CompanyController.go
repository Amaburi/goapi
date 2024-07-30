package CompanyController

import (
	"net/http"

	models "goapi/model/album"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var companies []models.Company
	if err := models.DB.Find(&companies).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

func Show(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := models.DB.First(&company, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Company not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"company": company})
}

func Create(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := models.DB.Create(&company).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"company": company})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Model(&company).Where("id = ?", id).Updates(&company)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company data successfully updated"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).Delete(&models.Company{}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company data deleted successfully"})
}
