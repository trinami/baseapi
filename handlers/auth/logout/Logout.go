package logout

import (
	"net/http"
	"encoding/base64"

	"api/db"
	"api/models"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	var logoutReq LogoutReq
    if err := c.ShouldBind(&logoutReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if len(logoutReq.Secret) > 44 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid"})
        return
	}

	byteArray, err := base64.StdEncoding.DecodeString(logoutReq.Secret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid"})
		return
	}

	result := db.DB.Where("secret = ?", byteArray).Delete(&models.Token{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown token"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Logout success"})
}
