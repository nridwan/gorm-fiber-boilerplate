package app

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AppModule struct {
	App             *fiber.App
	ResponseService ResponseService
}

func NewFiber(responseService ResponseService) *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: responseService.ErrorHandler,
	})
}

func NewModule(app *fiber.App, responseService ResponseService) *AppModule {
	return &AppModule{
		App:             app,
		ResponseService: responseService,
	}
}

func fxRegister(lifeCycle fx.Lifecycle, module *AppModule) {
	base.FxRegister(module, lifeCycle)
}

func SetupModule(config config.ConfigService) *AppModule {
	responseService := NewResponseService(config)
	return NewModule(NewFiber(responseService), responseService)
}

var FxModule = fx.Module("app", fx.Provide(NewFiber), fx.Provide(NewResponseService), fx.Provide(NewModule), fx.Invoke(fxRegister))

// implements `BaseModule` of `base/module.go` start

func (module *AppModule) OnStart() error {
	return nil
}

func (module *AppModule) OnStop() error {
	return nil
}

// implements `BaseModule` of `base/module.go` end
