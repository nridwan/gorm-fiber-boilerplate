package cmd

import (
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() error {
		<-c
		log.Println("stopping server ...")
		err := app.Shutdown()
		log.Println("stop server success")
		return err
	}()

	// ...

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}
}
