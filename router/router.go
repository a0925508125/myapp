package router

import (
	"myapp/router/controller"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	// Gin router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 處理 favicon.ico，避免 404 log
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No Content
	})
	return r
}

func SetupRouter(r *gin.Engine) {
	controllerCenter := controller.NewController()
	noLogin := r.Group("/v1")
	noLogin.GET("/hello", controllerCenter.GetHello)
}
