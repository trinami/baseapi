package models


func Register(c *gin.Context) {
	/*var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)*/
	var login Login

        // Binding the data from the request body to the struct
        if err := c.ShouldBind(&login); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		db.DB.Create(&user)
	c.JSON(http.StatusOK, "ok")
}
