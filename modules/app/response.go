package app

import (
	"gofiber-boilerplate/modules/app/appmodel"
	"gofiber-boilerplate/modules/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ResponseService interface {
	CreateErrorResponse(code int, message string, errors []appmodel.Error) *appmodel.Response
	CreateResponse(code int, message string, data interface{}) *appmodel.Response
	SendErrorResponse(ctx *fiber.Ctx, code int, message string, errors []appmodel.Error) error
	SendValidationErrorResponse(ctx *fiber.Ctx, code int, message string, errors validator.ValidationErrors) error
	SendResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error
	SendSuccessResponse(ctx *fiber.Ctx, message string, data interface{}) error
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

// impl `ResponseService` start

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

func (service *responseServiceImpl) SendErrorResponse(ctx *fiber.Ctx, code int, message string, errors []appmodel.Error) error {
	return ctx.Status(code).JSON(service.CreateErrorResponse(code, message, errors))
}

func (service *responseServiceImpl) SendValidationErrorResponse(ctx *fiber.Ctx, code int, message string, errors validator.ValidationErrors) error {
	mappedError := make([]appmodel.Error, len(errors))
	for i, err := range errors {
		mappedError[i] = appmodel.Error{
			Field:   err.Field(),
			Message: err.Error(),
		}
	}

	return service.SendErrorResponse(ctx, code, message, mappedError)
}

func (service *responseServiceImpl) SendResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	return ctx.Status(code).JSON(service.CreateResponse(code, message, data))
}

func (service *responseServiceImpl) SendSuccessResponse(ctx *fiber.Ctx, message string, data interface{}) error {
	return service.SendResponse(ctx, 200, message, data)
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

// impl `ResponseService` end
