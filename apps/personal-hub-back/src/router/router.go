package router

import (
	"github.com/gin-gonic/gin"
)

// New builds the HTTP engine with every route registered.
func New() *gin.Engine {
	g := gin.Default()
	serveRoutes(g)
	return g
}

func Run() {
	New().Run(":3000")
}
