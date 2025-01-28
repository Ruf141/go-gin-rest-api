package user

import "github.com/gin-gonic/gin"

type UserHandler struct{}

// GetUsers godoc
// @Summary ユーザー一覧を取得
// @Tags GetUsers
// @Accept json
// @Produce json
// @Success 200 {object} []ResponseUser
// @Router /v1/users [get]
func (h *UserHandler) GetUsers(ctx *gin.Context) {}

// GetUseById godoc
// @Summary ユーザー一覧を取得
// @Tags GetUseById
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} ResponseUser
// @Router /v1/users/:id [get]
func (h *UserHandler) GetUseById(ctx *gin.Context) {}

// GetUseById godoc
// @Summary ユーザー一覧を編集
// @Tags EdirUser
// @Accept json
// @Produce json
// @Param request body RequestUserParam  ture "ユーザー情報"
// @Success 200 {object} Response
// @Router /v1/users [post]
func (h *UserHandler) EdirUser(ctx *gin.Context) {}


