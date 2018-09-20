package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teddysingh/api_gin_gonic/db"
	"github.com/teddysingh/api_gin_gonic/handlers"
	"github.com/teddysingh/api_gin_gonic/utils"

	"github.com/teddysingh/api_gin_gonic/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbconn = db.Connect()
var logger = utils.Logger

func init() {
	logger.Println("Running migrations now...")
	dbconn.AutoMigrate(&models.User{})
}

func main() {
	defer dbconn.Close()

	serverDescription := map[string]string{
		"version": "1.0",
		"name":    "Hello World Server",
		"port":    ":3000",
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, serverDescription)
	})
	r.GET("/users", handlers.FindUser)
	r.POST("/users", handlers.CreateUser)

	logger.Infof("Starting %s on port %s\n", serverDescription["name"], serverDescription["port"])
	r.Run(serverDescription["port"])
}
