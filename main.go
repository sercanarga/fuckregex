package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Static("/assets", "web/assets")
	router.LoadHTMLGlob("web/templates/**/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	err := router.Run(":8181")
	if err != nil {
		panic(err)
	}
}
