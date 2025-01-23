package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health, systemHandler.Helth")
	}
	{
		userHandler := user.NewSystemHandler()
		v1.GET("", userHandler.GetUers)
		v1.GET("/:id, userHandler.GetUserById")
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
