// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"

	"github.com/adiazny/strava-service/cmd/services/strava-api/handlers/v1/testgrp"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	tgh := testgrp.Handlers{
		Log: cfg.Log,
	}

	mux.Handle(http.MethodGet, "/v1/test", tgh.Test)

	return mux
}
