package gateway

import (
	"github.com/gin-gonic/gin"
)

type Gateway struct {
	router *gin.Engine
}

func New() *Gateway {
	router := gin.Default()
	return &Gateway{router: router}
}

func (g *Gateway) Router() *gin.Engine {
	return g.router
}

func (g *Gateway) Run(addr string) error {
	return g.router.Run(addr)
}
