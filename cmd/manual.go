package cmd

import (
	"gofiber-boilerplate/base"
	appmodule "gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
)

func CommandManual() *cli.Command {
	return &cli.Command{
		Name:  "manual",
		Usage: "start manual server",
		Action: func(cCtx *cli.Context) error {
			runManual()
			return nil
		},
	}
}

func runManual() {
	configModule := config.SetupModule()
	appModule := appmodule.SetupModule(configModule)

	modules := []base.BaseModule{
		configModule,
		appModule,
	}

	for i := range modules {
		modules[i].OnStart()
	}

	appModule.App.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(400, "Error")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() error {
		<-c
		log.Println("stopping server ...")
		go func() {
			appModule.App.Shutdown()
		}()
		for i := range modules {
			modules[i].OnStop()
		}
		log.Println("stop server success")
		return nil
	}()

	// ...

	if err := appModule.App.Listen(configModule.Service.Getenv("APP_HOST", "") + ":" + configModule.Service.Getenv("PORT", "3000")); err != nil {
		log.Panic(err)
	}
}
