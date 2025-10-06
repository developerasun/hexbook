package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Show the health status
// @Description Get server health status
// @Tags health
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
// @Tags health
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
