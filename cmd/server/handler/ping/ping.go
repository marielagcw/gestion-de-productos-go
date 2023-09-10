package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func NewControllerPing() *Controller {
	return &Controller{}
}

// Ping godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (c *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
