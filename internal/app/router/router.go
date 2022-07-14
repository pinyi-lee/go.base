package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pinyi-lee/go.base.git/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pinyi-lee/go.base.git/internal/app/handler"
	"github.com/pinyi-lee/go.base.git/internal/pkg/config"
)

var Router *gin.Engine

func SetupRouter() (router *gin.Engine) {
	docs.SwaggerInfo.Version = config.Env.Version

	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()
	router.Use(handler.CORSMiddleware(), handler.ErrorMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/health", handler.HealthHandler)
	router.GET("/version", handler.VersionHandler)

	return
}

func Setup() error {
	Router = SetupRouter()

	return nil
}
