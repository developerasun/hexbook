package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

const (
	ROOT = ""
	PORT = 3010
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	hasError := env.Load("../.run.env")

	if hasError != nil {
		log.Fatalln("main.go: can't load secrets correctly", hasError.Error())
		return
	}

	root := router.Group(ROOT)

	root.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, gin.H{})
	})

	root.GET("/health", func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
