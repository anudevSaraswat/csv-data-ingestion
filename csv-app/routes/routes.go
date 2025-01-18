package routes

import (
	"csv-app/handler"

	middleware "csv-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(middleware.InitDatabase())
	engine.Use(gin.Logger())

	engine.GET("/user", handler.APIGetUser)

	return engine

}
