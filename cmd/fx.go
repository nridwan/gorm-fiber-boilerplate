package cmd

import (
	"context"
	appmodule "gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

func CommandFx() *cli.Command {
	return &cli.Command{
		Name:  "fx",
		Usage: "start fx server",
		Action: func(cCtx *cli.Context) error {
			runFx()
			return nil
		},
	}
}

func runFx() {
	app := fx.New(
		config.FxModule,
		appmodule.FxModule,
		fx.Invoke(RegisterWebServer),
	)

	app.Run()
}

func RegisterWebServer(
	lifeCycle fx.Lifecycle,
	app *fiber.App,
	configService *config.ConfigService,
) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				app.Get("/", func(c *fiber.Ctx) error {
					return fiber.NewError(400, "Error")
				})

				if err := app.Listen(
					configService.Getenv("APP_HOST", "") + ":" + configService.Getenv("PORT", "3000"),
				); err != nil {
					log.Fatalf("start server error : %v\n", err)
				}

			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("stopping server ...")
			err := app.Shutdown()
			log.Println("stop server success")
			return err
		},
	})
}
