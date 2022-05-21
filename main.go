package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fyyang/goctl-api2proto/action"
	"github.com/urfave/cli/v2"
)

var (
	version  = "20220521"
	commands = []*cli.Command{
		{
			Name:   "api2proto",
			Usage:  "generates proto files from api definition",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "filename",
					Usage: "proto save file name",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate proto files from api definition"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-api2proto: %+v\n", err)
	}
}
