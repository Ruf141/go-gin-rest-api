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

// GetUserById godoc
// @Summary ユーザーを取得
// @Tags GetUserById
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} ResponseUser
// @Router /v1/users/{id} [get]
func (h *UserHandler) GetUseById(ctx *gin.Context) {}

// EditUser godoc
// @Summary ユーザー情報を編集
// @Tags EditUser
// @Accept json
// @Produce json
// @Param request body RequestUserParam true "ユーザー情報"
// @Success 200 {object} Response
// @Router /v1/users [post]
func (h *UserHandler) EditUser(ctx *gin.Context) {}

// DeleteUser godoc
// @Summary ユーザーを削除
// @Tags DeleteUser
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} Response
// @Router /v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
