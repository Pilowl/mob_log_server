package main

import (
	"log"
	"os"
	"time"

	"./config"
	"./controllers"
	"./repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()

	if gin.Mode() == gin.ReleaseMode {
		log.SetFlags(0) // Disabling logs if isn't release
		config.Init("config/production.json")
	} else {
		// Initializing log output to file

		f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
		log.Print("Debug | Test Mode is turned ON, using default config.")
		config.Init("config/default.json")
	}

	// Initializing database
	repository.Init()

	// Initializing router
	r := gin.New()

	r.Use(gin.Logger(), cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test123",
		})
	})
	v1 := r.Group("/api/v1/log")
	{
		v1.POST("/create", controllers.AddNewLog)
		v1.POST("/append", controllers.AppendToLog)
		v1.GET("/sessions", controllers.GetSessions)
		v1.GET("/get/:id", controllers.GetLog) // Get log by session ID
	}

	r.Run(config.GetConfig().Port)

}
