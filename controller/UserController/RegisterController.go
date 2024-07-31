package UserController

import (
	"net/http"

	models "goapi/model/album"
	"goapi/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := models.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
