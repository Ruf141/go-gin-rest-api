package user

import"github.com/gin-gonic/gin"

type UserHandler struct {}

// GetUsers godoc
// @Summary ユーザー一覧を取得
// @Tags GetUsers
// @Accept json
// Produce json
// Success 200 {object} []ResponseUser
// Router /v1/users [get]
func(h *UserHandler) GetUsers(ctx *gin.Context){}


