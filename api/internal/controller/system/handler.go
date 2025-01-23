package system

import "github.com/gin-gonic/gin"

type SystemHandlre struct{}

// HealthCheck godoc
// @Summary 死活監視用
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (h *SystemHandlre) Helth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func NewSystemHandler() *SystemHandlre {
	return &SystemHandlre{}
}
