package presenter

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/o-ga09/api/internal/controller/system"
	"github.com/o-ga09/api/internal/controller/user"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)
	{
        systemHandler := system.NewSystemHandler()
        v1.GET("/health", systemHandler.Heath) 
	}
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers) 
		v1.GET("/:id", userHandler.GetUseById) 
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := r.Run() // ポートを指定しない場合はデフォルトの8080
	if err != nil {
		return err
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
