package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	pkg "github.com/developerasun/hexbook/pkg"
	"github.com/gin-gonic/gin"
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
		ctx.Error(err)
		return
	}
	log.Println("qrcodeData: ", qrcodeData)

	wallet := ctx.PostForm("wallet")
	amount := ctx.PostForm("amount")
	appType := ctx.PostForm("apptype")
	tokenType := ctx.PostForm("tokentype")

	log.Println("app type: ", appType)
	log.Println("amount: ", amount)
	log.Println("tokentype: ", tokenType)

	if len(wallet) == 0 {
		ctx.Error(errors.New("RenderQrCode:len(wallet): empty wallet from client"))
		return
	}

	filename, err := pkg.GenerateQrCode(appType, wallet, amount, tokenType)

	var _html string
	if err != nil {
		_html = fmt.Sprintf(`<div class="text-error">%s</div>`, err.Error())
	} else {
		_html = fmt.Sprintf(`<div><img src="%s" alt="qrcode"/></div>`, "/assets/qrcode/"+filename)
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(_html))
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
func RenderQrCode2(ctx *gin.Context) {

	var qrcodeData QRCodeDataDto2

	if err := ctx.ShouldBind(&qrcodeData); err != nil {
		ctx.Error(err)
		return
	}
	log.Println("qrcodeData2: ", qrcodeData)

	wallet := ctx.PostForm("wallet2")
	amount := ctx.PostForm("amount2")
	appType := ctx.PostForm("apptype2")
	tokenType := ctx.PostForm("tokentype2")

	log.Println("apptype2: ", appType)
	log.Println("amount2: ", amount)
	log.Println("tokentype2: ", tokenType)

	if len(wallet) == 0 {
		ctx.Error(errors.New("RenderQrCode2:len(wallet): empty wallet from client"))
		return
	}

	filename, err := pkg.GenerateQrCode(appType, wallet, amount, tokenType)

	var _html string
	if err != nil {
		_html = fmt.Sprintf(`<div class="text-error">%s</div`, err.Error())
	} else {
		_html = fmt.Sprintf(`<div><img src="%s" alt="qrcode"/></div>`, "/assets/qrcode/"+filename)
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(_html))
}
