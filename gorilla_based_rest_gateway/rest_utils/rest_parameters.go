package rest_utils

import (
"github.com/codegangsta/cli"
)

var (
	RestGatewayAddress = cli.StringFlag{
		Name:  "rest_server_address_port",
		Usage: "<ip>:<port> to listen on",
		Value: "127.0.0.1:8101",
	}
)