package main

import (
	"log"
	"os"
	"strings"

	apiController "github.com/fatcat/internal/api"
	"github.com/fatcat/internal/constant"
	"github.com/fatcat/internal/database"
	"github.com/gin-gonic/gin"

	docs "github.com/fatcat/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	env "github.com/joho/godotenv"
)

// @title Fatcat API
// @version 1.0
// @description Fatcat backend API documentation
// @BasePath /
func main() {
	dir, gErr := os.Getwd()

	if gErr != nil {
		log.Fatalln(gErr.Error())
	}

	envPath := strings.Join([]string{dir, "/", ".run.env"}, "")
	log.Println("main.go: envPath: " + envPath)

	hasError := env.Load(envPath)
	if hasError != nil {
		log.Fatalln("main.go: can't load secrets correctly", hasError.Error())
		return
	}
	log.Println("main.go: env loaded")

	schema.ConnectAndMigrate()
	log.Println("main.go: database connected")

	log.Println("main.go: start initiating gin server")
	router := gin.Default()
	router.SetTrustedProxies(nil)

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	root := router.Group(constant.ROUTE_ROOT)
	root.GET("/", apiController.RenderMainPage)

	api := router.Group(constant.ROUTE_API)
	api.GET("/health", apiController.Health)
	api.GET("/fetch", apiController.FetchDummyData)
	api.GET("/ws", apiController.ConnectWebsocket)
	// @dev htmx test
	api.GET("/clicked", apiController.RenderClicked)

	router.Run(":" + os.Getenv("PORT"))
	log.Println("main.go: router started")
}
