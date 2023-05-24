package routes

import (
	"Granary/hander"
	"Granary/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	engine := gin.Default()
	engine.Use(middleware.Cors())
	auth := engine.Group("api/v1")
	{

		auth.GET("/get_all", hander.GetAll)
	}
	return engine
}
