package config

import (
	"gofiber-boilerplate/base"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type ConfigModule struct {
	Service *ConfigService
}

func NewModule(service *ConfigService) *ConfigModule {
	return &ConfigModule{
		Service: service,
	}
}

func fxRegister(lifeCycle fx.Lifecycle, module *ConfigModule) {
	base.FxRegister(module, lifeCycle)
}

func SetupModule() *ConfigModule {
	return NewModule(&ConfigService{})
}

var FxModule = fx.Module("Config", fx.Provide(NewService), fx.Provide(NewModule), fx.Invoke(fxRegister))

// implements `BaseModule` of `base/module.go` start

func (module *ConfigModule) OnStart() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return nil
}

func (module *ConfigModule) OnStop() error {
	return nil
}

// implements `BaseModule` of `base/module.go` end
