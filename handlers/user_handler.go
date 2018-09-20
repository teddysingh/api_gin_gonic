package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddysingh/api_gin_gonic/utils"

	"github.com/teddysingh/api_gin_gonic/db"
	"github.com/teddysingh/api_gin_gonic/models"
)

var logger = utils.Logger

// CreateUser - Handles POST action
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	logger.Debugf("User is: %v", user)

	dbconn := db.DB
	if dbconn != nil {
		if err := dbconn.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"Message": "Error while persisting user",
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		logger.Errorf("DB Connection State = %v\n", dbconn)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error while connecting to DB",
		})
	}
}

// FindUser - Get user by ID
func FindUser(c *gin.Context) {
	var user models.User
	queryEmail := c.Query("email")
	dbconn := db.DB
	if dbconn != nil {
		if err := dbconn.Where("email LIKE ?", "%"+queryEmail+"%").Find(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, map[string]string{
				"Message": "Error while finding user",
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		logger.Errorf("DB Connection State = %v\n", dbconn)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": "Error while connecting with DB",
		})
	}
}
