package main

import (
	"github.com/gin-gonic/gin"
	apiFunc "github.com/Shashank-1993/completeGoProj/api"

	_ "github.com/Shashank-1993/completeGoProj/docs" // This is important for swagger to recognize the docs
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main(){
	r := gin.Default();

	// Add the Swagger documentation route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("api/v1")

	api.GET("ready", apiFunc.Ready)

	r.Run(":8080")
}