package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	router *gin.Engine
}

func NewServer(router *gin.Engine) Server {
	return Server{"8080", router}
}

func (s *Server) Run() {
	log.Println("server is running at port: ", s.port)
	log.Fatal(s.router.Run(":" + s.port))
}
