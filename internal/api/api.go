package api

import (
	"fmt"
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

// RenderMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags view
// @Router / [get]
func RenderMainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

// RenderQrCode godoc
// @Summary request a corresponding qrcode for the submitted wallet
// @Description request a corresponding qrcode for the submitted wallet
// @Tags api
// @Accept application/x-www-form-urlencoded
// @Produce html
// @Param wallet formData string true "Wallet address" default(0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11)
// @Success 200 {string} string <div><img src="image-path" alt="qrcode"/></div>
// @Router /api/qrcode [post]
func RenderQrCode(ctx *gin.Context) {

	var qrcodeData QRCodeDataDto

	if err := ctx.ShouldBind(&qrcodeData); err != nil {
		log.Fatalln("RenderQrCode: ", err.Error())
	}
	log.Println("qrcodeData: ", qrcodeData)

	wallet := ctx.PostForm("wallet")
	amount := ctx.PostForm("amount")
	appType := ctx.PostForm("apptype")

	log.Println("app type: ", appType)
	log.Println("amount: ", amount)

	if len(wallet) == 0 {
		log.Fatalln("RenderQrCode:len(wallet): empty wallet from client")
	}

	filename := pkg.GenerateQrCode(appType, wallet, amount)

	_html := fmt.Sprintf(`<div><img src="%s" alt="qrcode"/></div>`, "/assets/qrcode/"+filename)
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(_html))
}
