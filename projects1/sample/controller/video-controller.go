package controller

import (
	"github/silasgah/news/entity"
	"github/silasgah/news/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func customFunc(fl validator.FieldLevel) bool {
	if fl.Field().String() == "Cool" {
		return false
	}

	return true
}

// The function that export the controller
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", customFunc)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
