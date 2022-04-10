// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	"github.com/adiazny/strava-service/cmd/services/strava-api/handlers/v1/testgrp"
	"github.com/adiazny/strava-service/foundation/web"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {

	// Construct the web.App which holds all routes.
	app := web.NewApp(cfg.Shutdown)

	// Load the routes for the different versions of the API
	v1(app, cfg)

	return app
}

// v1 binds all the version 1 routes.
func v1(app *web.App, cfg APIMuxConfig) {
	const version = "v1"

	tgh := testgrp.Handlers{
		Log: cfg.Log,
	}

	app.Handle(http.MethodGet, version, "/test", tgh.Test)
}
