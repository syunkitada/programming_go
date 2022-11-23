// This file is safe to edit. Once it exists it will not be overwritten

package api

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/syunkitada/programming_go/app-sample/pkg/lib/api_middleware"
	"github.com/syunkitada/programming_go/app-sample/pkg/lib/api_runtime"
	lib_config "github.com/syunkitada/programming_go/app-sample/pkg/lib/config"
	"github.com/syunkitada/programming_go/app-sample/pkg/lib/logger"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/app"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/config"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/gen/api/operations"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/gen/api/operations/todos"
	"github.com/syunkitada/programming_go/app-sample/pkg/sample-api/gen/models"
)

//go:generate swagger generate server --target ../../../../../app-sample --name SampleAPI --spec ../../swagger.yml --model-package pkg/sample-api/gen/models --server-package pkg/sample-api/gen/api --principal interface{}

var appFlags = struct {
	ConfigFiles []string `short:"c" long:"config-files" description:"config files"`
}{}

func configureFlags(api *operations.SampleAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "App Flags",
			LongDescription:  "",
			Options:          &appFlags,
		},
	}
}

func configureAPI(api *operations.SampleAPIAPI) http.Handler {
	var conf config.Config
	lib_config.MustLoadConfigFiles(&conf, &config.DefaultConfig, appFlags.ConfigFiles)
	logger.Init(&conf.Logger)
	app := app.New(&conf)

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = api_runtime.JSONConsumer()
	api.JSONProducer = api_runtime.JSONProducer()

	api.TodosAddOneHandler = todos.AddOneHandlerFunc(func(params todos.AddOneParams) middleware.Responder {
		if err := app.AddItem(params.Body); err != nil {
			return todos.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return todos.NewAddOneCreated().WithPayload(params.Body)
	})
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(func(params todos.DestroyOneParams) middleware.Responder {
		if err := app.DeleteItem(params.ID); err != nil {
			return todos.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return todos.NewDestroyOneNoContent()
	})
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
		mergedParams := todos.NewFindTodosParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return todos.NewFindTodosOK().WithPayload(app.AllItems(*mergedParams.Since, *mergedParams.Limit))
	})
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(func(params todos.UpdateOneParams) middleware.Responder {
		if err := app.UpdateItem(params.ID, params.Body); err != nil {
			return todos.NewUpdateOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return todos.NewUpdateOneOK().WithPayload(params.Body)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return api_middleware.CommonHandler(handler)
}
