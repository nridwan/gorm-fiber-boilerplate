package user

import (
	"gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/app/appmodel"
	"gofiber-boilerplate/modules/user/userdto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	validationError = "Validation Error"
)

type userController struct {
	service         UserService
	responseService app.ResponseService
	validator       *validator.Validate
}

func newUserController(service UserService, responseService app.ResponseService, validator *validator.Validate) *userController {
	return &userController{
		service:         service,
		responseService: responseService,
		validator:       validator,
	}
}

// handlers start

func (controller *userController) handleCreate(ctx *fiber.Ctx) error {
	request := userdto.CreateUserDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	model, err := controller.service.Insert(request.ToModel())

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendResponse(ctx, 201, "Success", model)
}

func (controller *userController) handleList(ctx *fiber.Ctx) error {
	request := appmodel.NewGetListRequest(ctx.Query("page"), ctx.Query("limit"), ctx.Query("search"))
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	list, err := controller.service.List(request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, "Success", appmodel.PaginationResponse{
		List: &appmodel.PaginationResponseList{
			Pagination: &appmodel.PaginationResponsePagination{},
			Content:    list,
		},
	})
}

func (controller *userController) handleDetail(ctx *fiber.Ctx) error {
	user, err := controller.service.Detail(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return controller.responseService.SendSuccessResponse(ctx, "Success", user)
}

func (controller *userController) handleUpdate(ctx *fiber.Ctx) error {
	request := userdto.UpdateUserDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	user, err := controller.service.Update(ctx.Params("id"), &request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return controller.responseService.SendSuccessResponse(ctx, "Success", user)
}

func (controller *userController) handleDelete(ctx *fiber.Ctx) error {
	err := controller.service.Delete(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return controller.responseService.SendSuccessResponse(ctx, "Success", nil)
}

// handlers end
