package controller

import (
	"GOREST/entity"
	"GOREST/service"
	"GOREST/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)

	if err != nil {
		return err
	}
	
	controller.service.Save(video)
	return nil
}
