package app

import (
	"gofiber-boilerplate/modules/app/appmodel"
	"gofiber-boilerplate/modules/config"

	"github.com/gofiber/fiber/v2"
)

type ResponseService interface {
	CreateErrorResponse(code int, message string, errors []appmodel.Error) *appmodel.Response
	CreateResponse(code int, message string, data interface{}) *appmodel.Response
	ErrorHandler(ctx *fiber.Ctx, err error) error
}

type responseServiceImpl struct {
	appName string
}

func NewResponseService(config config.ConfigService) ResponseService {
	return &responseServiceImpl{
		appName: config.Getenv("APP_CODE", "APP"),
	}
}

func (service *responseServiceImpl) CreateErrorResponse(code int, message string, errors []appmodel.Error) *appmodel.Response {
	return &appmodel.Response{
		ResponseSchema: &appmodel.ResponseSchema{
			ResponseCode:    &code,
			ResponseMessage: &message,
		},
		ResponseOutput: appmodel.ErrorResponse{
			Errors: errors,
		},
	}
}

func (service *responseServiceImpl) CreateResponse(code int, message string, data interface{}) *appmodel.Response {
	return &appmodel.Response{
		ResponseSchema: &appmodel.ResponseSchema{
			ResponseCode:    &code,
			ResponseMessage: &message,
		},
		ResponseOutput: data,
	}
}

// ErrorHandler check if connection should be continued or not
func (service *responseServiceImpl) ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(service.CreateResponse(code, err.Error(), nil))
}
