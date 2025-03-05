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

// Ping godoc
//
//	@Summary		Ping server
//	@Description	Check if the server is running
//	@Tags			health
//	@Produce		plain
//	@Success		200	{string}	string	"pong"
//	@Router			/api/v1/ping [get]
func (c *PingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
