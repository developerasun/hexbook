package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatcat/auth"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	env "github.com/joho/godotenv"
)

const (
	ROOT = ""
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

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	hasError := env.Load("../.run.env")

	// @dev init connection manager instance on startup
	sc := &auth.SocketManager{
		List:      make(map[int]*auth.ClientContext),
		Count:     0,
		MaxClient: 500,
	}

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

	router.Run(":" + os.Getenv("PORT"))
}
