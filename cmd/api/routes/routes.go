package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "myApi/docs"
	"myApi/pkg/profile"
)

func Configure(router *gin.Engine) {

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.POST("/api/v1/profiles", JSON(profile.Create))
}
