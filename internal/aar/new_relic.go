package aar

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
	"os"
)

const appName = "AAR"

var (
	newRelic *newrelic.Application
)

// SetupNewRelic sets up monitoring in NewRelic
func SetupNewRelic(licenseKey string, r *mux.Router) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting New Relic: %q", err)
		os.Exit(1)
	}

	newRelic = app
	r.Use(nrgorilla.Middleware(app))
}
