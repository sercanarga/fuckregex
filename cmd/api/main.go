package main

import (
	"fuckregex/handler"
	"fuckregex/internal"
	"fuckregex/internal/db"
	"fuckregex/middleware"
	"fuckregex/model/db_model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	// Setup logger
	internal.SetupLogger()

	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	err = db.ConnectDB(&db_model.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	// Rate limiter middleware
	middleware.RateLimitMiddleware(app)

	app.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})

	//Handlers
	app.GET("/latest", handler.Latest)

	app.POST("/generate", handler.Generate)
	app.POST("/get", handler.Get)
	app.POST("/report", handler.Report)

	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
