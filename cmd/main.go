package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	api "github.com/fatcat/internal/api"
	"github.com/fatcat/internal/constant"
	"github.com/fatcat/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	docs "github.com/fatcat/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	env "github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		whitelist := os.Getenv("WHITELIST")
		log.Println(r.Host, whitelist, "is whitelisted: ", r.Host == whitelist)
		return r.Host == whitelist
	},
}

// @title Fatcat API
// @version 1.0
// @description Fatcat backend API documentation
// @BasePath /
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

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

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	root := router.Group(constant.ROUTE_ROOT)
	root.GET("/", api.ViewMainPage)
	root.GET("/health", api.Health)
	root.GET("/fetch", api.FetchDummyData)
	root.GET("/ws", api.ConnectWebsocket)

	// @dev htmx test
	router.GET("/clicked", api.ViewClicked)

	router.Run(":" + os.Getenv("PORT"))
	log.Println("main.go: router started")
}
