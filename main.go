package main

import (
	"fmt"
	"gofiber-boilerplate/cmd"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	commands := []*cli.Command{
		cmd.CommandManual(),
		cmd.CommandFx(),
	}

	app := &cli.App{
		Commands: commands,
		Name:     "apiserver",
		Usage:    "manual, fx",
		Action: func(cli *cli.Context) error {
			fmt.Printf("%s version:%s\n", cli.App.Name, "3.0")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
