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

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	SetupLogOutput()
	//Middlewares
	server.Use(gin.Recovery(),
		middleware.Logger(),
		middleware.Authorization(),
	)

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})

		apiRoutes.GET("/video/all", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/video", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	PORT := os.Getenv("PORT")

	if PORT =="" {
		PORT = "8080"
	} 
	
	server.Run(":"+PORT)
}
