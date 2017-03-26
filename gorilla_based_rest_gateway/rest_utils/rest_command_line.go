package rest_utils

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	"github.com/fsgura/testing-go/gorilla_based_rest_gateway/rest_version"
)

func Run(name string, commands ... cli.Command) {
	app := cli.NewApp()
	app.Name = name

	versionCommand := cli.Command{
		Name:      "version",
		ShortName: "v",
		Usage:     "Show the version",
		Action: func(_ *cli.Context) {
			fmt.Printf("%s-%s\n", rest_version.VERSION, rest_version.FEATURE)
		},
	}

	app.Commands = append(commands, versionCommand)
	app.Version = fmt.Sprintf("%s-%s", rest_version.VERSION, rest_version.FEATURE)
	app.Run(os.Args)
}
