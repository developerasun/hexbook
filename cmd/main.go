package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	apiController "github.com/hexbook/internal/api"
	"github.com/hexbook/internal/constant"

	docs "github.com/hexbook/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	env "github.com/joho/godotenv"
)

// @title hexbook API
// @version 1.0
// @description hexbook backend API documentation
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
	api.POST("/qrcode", apiController.RenderQrCode)

	router.Run(":" + os.Getenv("PORT"))
	log.Println("main.go: router started")
}
