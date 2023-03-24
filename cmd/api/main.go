package main

import (
	"fuckregex/internal"
	"fuckregex/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// Setup logger
	internal.SetupLogger()

	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	app := gin.Default()

	// Rate limiter middleware
	middleware.RateLimitMiddleware(app)

	app.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	//Handlers
	//app.POST("/login", handler.Login)

	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
