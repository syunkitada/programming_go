package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/syunkitada/programming_go/project_examples/shop/api/shop-api/oapi"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/config"
	"github.com/syunkitada/programming_go/project_examples/shop/internal/shop-api/handler"
)

func main() {
	conf := config.GetDefaultConfig()

	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	apiHandler := handler.NewHandler(&conf)

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	oapi.RegisterHandlers(e, apiHandler)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
