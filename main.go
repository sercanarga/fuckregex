package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
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
