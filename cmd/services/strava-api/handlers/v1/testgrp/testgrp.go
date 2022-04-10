package testgrp

import (
	"context"
	"net/http"

	"github.com/adiazny/strava-service/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of test endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler for my development.
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "ok",
	}

	statusCode := http.StatusOK
	h.Log.Infow("test", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

	return web.Respond(ctx, w, status, http.StatusOK)
}
