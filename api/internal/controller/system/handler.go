package system

import "github.com/gin-gonic/gin"

type SystemHandlre struct{}

func (h *SystemHandlre) Helth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func NewSystemHandler() *SystemHandlre {
	return &SystemHandlre{}
}
