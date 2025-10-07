package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fatcat/internal/auth"
	"github.com/fatcat/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Health godoc
// @Summary Show the health status
// @Description Get server health status
// @Tags dev
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func Health(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// FetchDummyData godoc
// @Summary Fetch data from jsonplaceholder
// @Description Fetch data from jsonplaceholder
// @Tags dev
// @Produce json
// @Success 200 {object} map[string]string
// @Router /fetch [get]
func FetchDummyData(ctx *gin.Context) {
	resp, gErr := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if gErr != nil {
		log.Fatalln(gErr.Error())
	}
	defer resp.Body.Close()

	body, rErr := io.ReadAll(resp.Body)

	if rErr != nil {
		log.Fatalln(rErr.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": json.RawMessage(body),
	})
}

// ConnectWebsocket godoc
// @Summary upgrade header from http to ws
// @Description upgrade header from http to ws
// @Tags api
// @Router /ws [get]
func ConnectWebsocket(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			whitelist := os.Getenv("WHITELIST")
			log.Println(r.Host, whitelist, "is whitelisted: ", r.Host == whitelist)
			return r.Host == whitelist
		},
	}

	// @dev init connection manager instance on startup
	sc := &auth.SocketManager{
		List:      make(map[int]*auth.ClientContext),
		Count:     0,
		MaxClient: constant.MAX_SOCKET_CLIENT,
	}
	log.Println("api.go: socket manager initialized")

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
}

// ViewMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags api
// @Router / [get]
func ViewMainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main Website",
		"data": []struct {
			Name string
			Age  uint
		}{
			{Name: "jake", Age: 31},
			{Name: "brian", Age: 22},
			{Name: "smith", Age: 14},
		},
	})
}

// ViewClicked godoc
// @Summary testing htmx get method with swapping response html
// @Description testing htmx get method with swapping response html
// @Tags api
// @Router /clicked [get]
func ViewClicked(ctx *gin.Context) {
	_html := "<div>hello htmx there</div>"
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(_html))
}
