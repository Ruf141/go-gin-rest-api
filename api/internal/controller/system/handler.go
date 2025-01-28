package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}

// Response 通常のレスポンス
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// HealthCheck godoc
// @Summary 死活監視用
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func (h *SystemHandler) Heath(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
