package login

import (
	"baseapi/db"
	"baseapi/models"
	"baseapi/utils/validation"
	"baseapi/utils"
	
	"net/http"
	"encoding/base64"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq LoginReq
    if err := c.ShouldBind(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	valid, err := validation.ValidateUsername(loginReq.Username)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err = validation.ValidatePassword(loginReq.Password)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	valid, err = utils.CompareHashAndPassword(user.Password, loginReq.Password)
	if !valid || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wrong password"})
		return
	}

	var token *models.Token = &models.Token{}
	err = token.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	encodedToken := base64.StdEncoding.EncodeToString(token.Secret)

	user.Token = *token
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   encodedToken,
	})
}
