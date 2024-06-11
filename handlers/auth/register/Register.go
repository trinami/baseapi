package register

import (
	"net/http"
	"encoding/base64"
	"time"

	"baseapi/db"
	"baseapi/models"
	"baseapi/utils"
	"baseapi/utils/validation"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerReq RegisterReq
    if err := c.ShouldBind(&registerReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if registerReq.Password != registerReq.RepeatPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords didn't match"})
		return
	}

	valid, err := validation.ValidateUsername(registerReq.Username)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err = validation.ValidatePassword(registerReq.Password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("username = ?", registerReq.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := utils.GenerateHash(registerReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var token *models.Token = &models.Token{}
	err = token.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	var user models.User
	user.Username = registerReq.Username
	user.Password = hashedPassword
	user.Token = *token
	user.LastSeen = time.Now()

	encodedToken := base64.StdEncoding.EncodeToString(token.Secret)

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
		"token":   encodedToken,
	})
}
