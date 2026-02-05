package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saf1o/go-test/internal/handler"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/health", handler.Health)
	return r
}
