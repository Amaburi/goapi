package UserController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	models "goapi/model/album"
	"goapi/utils"
)

// Login handles user login and token generation
func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
