package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Show the health status222
// @Description Get server health status2222222
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
