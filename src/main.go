package main

import (
	apiFunc "github.com/Shashank-1993/completeGoProj/api"
	"github.com/gin-gonic/gin"

	_ "github.com/Shashank-1993/completeGoProj/docs" // This is important for swagger to recognize the docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Integration API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @basePath /api/v1
// @schemes http https

func main() {
	r := gin.Default()
	api := r.Group("api/v1")
	{
		api.GET("/ready", apiFunc.Ready)
	}
	// Add the Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
