package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/xqk/god-swagger/action"
	"os"
	"runtime"
)

var (
	version  = "20220314"
	commands = []*cli.Command{
		{
			Name:   "swagger",
			Usage:  "generates swagger.json",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Usage: "api request address",
				},
				&cli.StringFlag{
					Name:  "basepath",
					Usage: "url request prefix",
				},
				&cli.StringFlag{
					Name:  "filename",
					Usage: "swagger save file name",
				},
			},
		},
	}
)

func main() {
	// 调试：echo '{"apiFilePath":"./example/project.api","style":"","dir":"./example"}' | go run main.go swagger --filename project.json
	// 编译后：god api plugin -plugin god-swagger="swagger -filename project.json" -api ./example/project.api -dir ./example
	app := cli.NewApp()
	app.Usage = "a plugin of god to generate swagger.json"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("god-swagger: %+v\n", err)
	}
}
