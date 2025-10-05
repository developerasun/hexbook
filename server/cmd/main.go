package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatcat/internal/auth"
	"github.com/fatcat/internal/constant"
	"github.com/fatcat/internal/controller"
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

	hasError := env.Load("../.run.env")
	if hasError != nil {
		log.Fatalln("main.go: can't load secrets correctly", hasError.Error())
		return
	}
	log.Println("main.go: env loaded")

	schema.ConnectAndMigrate()
	log.Println("main.go: database connected")

	// @dev init connection manager instance on startup
	sc := &auth.SocketManager{
		List:      make(map[int]*auth.ClientContext),
		Count:     0,
		MaxClient: constant.MAX_SOCKET_CLIENT,
	}
	log.Println("main.go: socket manager initialized")

	root := router.Group(constant.ROUTE_ROOT)

	root.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNoContent, gin.H{})
	})

	// TODO replace to controller/service/repository dir
	root.GET("/health", controller.Health)
	root.GET("/fetch", controller.FetchDummyData)
	root.GET("/ws", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		result := make(chan *websocket.Conn)

		go func() {
			hErr := auth.HandleConnection(sc, conn)
			if hErr != nil {
				log.Println(hErr.Error())
				ctx.AbortWithError(http.StatusInternalServerError, hErr)
				return
			}
			result <- conn
		}()
		connection := <-result

		mt, payload, rErr := connection.ReadMessage()
		if rErr != nil {
			log.Println(rErr.Error())
			ctx.AbortWithError(http.StatusInternalServerError, rErr)
			return
		}
		client := sc.List[sc.Count]

		message := append(payload, []byte("with id: "+fmt.Sprint(client.SocketID))...)
		log.Println("socket id: ", client.SocketID)
		wErr := connection.WriteMessage(mt, message)

		if wErr != nil {
			log.Println(wErr.Error())
			ctx.AbortWithError(http.StatusInternalServerError, wErr)
			return
		}
	})

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":" + os.Getenv("PORT"))
	log.Println("main.go: router started")
}
