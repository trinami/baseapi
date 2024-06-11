package handlers

import (
	"baseapi/db"
	"baseapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetChallenge(c *gin.Context) {
	//create sha256 challenge & save to db
	c.JSON(http.StatusOK, "ok")
}
