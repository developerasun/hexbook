package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	env "github.com/joho/godotenv"
)

const (
	ROOT = ""
	PORT = 3010
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

type SocketConnection struct {
	list  map[int]*websocket.Conn
	count uint
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	hasError := env.Load("../.run.env")
	sc := &SocketConnection{
		list:  make(map[int]*websocket.Conn),
		count: 0,
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

		var rwMutex sync.RWMutex
		rwMutex.Lock()
		sc.list[int(sc.count)] = conn
		sc.count++
		rwMutex.Unlock()

		log.Println("mutex unlocked and socket connected")

		// for {
		messageType, payload, rErr := conn.ReadMessage()

		if rErr != nil {
			log.Println(rErr.Error())
			ctx.AbortWithError(http.StatusInternalServerError, rErr)
			return
		}

		customMessage := []byte(fmt.Sprintf(": woof woof: %v", sc.count))

		wErr := conn.WriteMessage(messageType, append(payload, customMessage...))
		if wErr != nil {
			log.Println(wErr.Error())
			ctx.AbortWithError(http.StatusInternalServerError, wErr)
			return
		}

		secondMessage := []byte(fmt.Sprintf(": meow meow: %v", sc.count))
		time.Sleep(5 * time.Second)

		conn.WriteMessage(messageType, secondMessage)
	})

	router.Run(":" + os.Getenv("PORT"))
}
