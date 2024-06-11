package antispam

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChallenge(c *gin.Context) {
	//create sha256 challenge & save to db
	c.JSON(http.StatusOK, "ok")
}
