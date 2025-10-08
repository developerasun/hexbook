package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pkg "github.com/hexbook/pkg"
)

// Health godoc
// @Summary Show the health status
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /api/health [get]
func Health(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, HealthResponse{
		Message: "ok",
	})
}

// FetchDummyData godoc
// @Summary Fetch data from jsonplaceholder
// @Description Fetch data from jsonplaceholder
// @Tags api
// @Produce json
// @Success 200 {object} FetchDummyDataResponse
// @Router /api/fetch [get]
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

	ctx.JSON(http.StatusOK, FetchDummyDataResponse{
		Data: json.RawMessage(body),
	})
}

// RenderMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags view
// @Router / [get]
func RenderMainPage(ctx *gin.Context) {
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

// RenderClicked godoc
// @Summary testing htmx get method with swapping response html
// @Description testing htmx get method with swapping response html
// @Tags api
// @Router /api/qrcode [get]
func RenderQrCode(ctx *gin.Context) {
	wallet := ctx.PostForm("wallet")

	if len(wallet) == 0 {
		log.Fatalln("RenderQrCode:len(wallet): empty wallet from client")
	}

	filename := pkg.GenerateQrCode(wallet)

	_html := fmt.Sprintf(`<div><img src="%s" alt="qrcode"/></div>`, "/assets/qrcode/"+filename)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(_html))
}
