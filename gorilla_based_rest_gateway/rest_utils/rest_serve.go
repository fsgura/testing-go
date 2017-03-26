package rest_utils

import (
	"github.com/codegangsta/cli"

	"net/http"
	"golang.org/x/net/context"
)

func ServeCmd(c *cli.Context, ctx context.Context, routes map[string]map[string]Handler) error {
	r := NewRouter(ctx, routes)
	return http.ListenAndServe(c.String("rest_server_address_port"), r)
}

