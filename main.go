package main

import (
	"GOREST/controller"
	"GOREST/middleware"
	"GOREST/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func SetupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	SetupLogOutput()
	//Middlewares
	server.Use(gin.Recovery(),
		middleware.Logger(),
		middleware.Authorization(),
	)

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/video/all", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			 ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
		}
	})

	server.Run(":8080")
}
