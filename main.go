package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hari0205/spotbuzz-task/controllers"
	"github.com/hari0205/spotbuzz-task/setup"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	// Initializing logger
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	// Setup DB
	setup.SetUpDB()

	// Middlewares
	router.Use(cors.Default())

	router.Use(gin.Recovery())

	apiGrp := router.Group("/api")
	{
		// TEST ROUTE
		apiGrp.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "PONG")
		})

		// PLAYER ROUTES
		apiGrp.POST("/players", controllers.CreatePlayer)
		apiGrp.GET("/players", controllers.GetPlayers)
		apiGrp.PUT("/players/:id", controllers.UpdatePlayer)
		apiGrp.GET("/players/random", controllers.GetRandomPlayers)
		apiGrp.GET("/players/:val", controllers.GetPlayerByRank)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Please check the endpoint and try again.",
		})
	})

	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":  http.StatusMethodNotAllowed,
			"error": "Please check the HTTP method and try again.",
		})
	})
	router.Run(":8080") // Firewall warning bypass

}
