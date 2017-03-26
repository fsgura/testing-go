/*
   Program     : rest_server.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This is the first example of a gorilla
                 based RESTful API gateway. This file is
                 the main package, from which the listening
                 http service spins up.
   Site        : http://www.github.com/fsgura/testing-go
 */
package main

import (
	"golang.org/x/net/context"
	"github.com/codegangsta/cli"
	"github.com/fsgura/testing-go/gorilla_based_rest_gateway/rest_utils"
	"os"
)

var routes = map[string]map[string]rest_utils.Handler {
	"POST": {
		"/v1/crowd/{id}": rest_utils.CreatePerson,
	},
	"GET": {
		"/v1/crowd": rest_utils.GetCrowd,
		"/v1/crowd/{id}": rest_utils.GetPersonById,
	},
	"DELETE": {
		"/v1/crowd/{id}": rest_utils.DeletePerson,
	},
}

func main() {
	serveCommand := cli.Command {
		Name:      "serve",
		ShortName: "s",
		Usage:     "Serve the API",
		Flags:     []cli.Flag{rest_utils.RestGatewayAddress},
		Action:    action(serveAction),
	}

	rest_utils.Run("gorilla_based_rest_gateway", serveCommand)
}

func serveAction(c *cli.Context) error {
	ctx := context.Background()
	return rest_utils.ServeCmd(c, ctx, routes)
}

func action(f func(c *cli.Context) error) func(c *cli.Context) {
	return func(c *cli.Context) {
		err := f(c)
		if err != nil {
			os.Exit(1)
		}
	}
}


