package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (c *PingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
