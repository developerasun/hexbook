package main

import (
	"fmt"
	"log"
	"net/http"
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
	wd, gErr := os.Getwd()

	if gErr != nil {
		log.Fatalln(gErr.Error())
	}

	envPath := strings.Join([]string{wd, "/", ".run.env"}, "")
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

	staticPath := strings.Join([]string{wd, "assets"}, "/")

	router.Static("/assets", staticPath)
	router.LoadHTMLGlob("templates/*")
	router.Use(ErrorHandler())
	router.Use(gin.Recovery())

	root := router.Group(constant.ROUTE_ROOT)
	root.GET("/", apiController.RenderMainPage)

	api := router.Group(constant.ROUTE_API)
	api.GET("/health", apiController.Health)
	api.POST("/qrcode", apiController.RenderQrCode)
	api.POST("/qrcode2", apiController.RenderQrCode2)

	router.Run(":" + os.Getenv("PORT"))
	log.Println("main.go: router started")
}

// ErrorHandler captures errors and returns a consistent JSON error response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			_html := fmt.Sprintf(`<div class="text-error">%s</div>`, err.Error())
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte(_html))
		}

		// Any other steps if no errors are found
	}
}
