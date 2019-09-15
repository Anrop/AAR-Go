package aar

import (
	"fmt"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

const appName = "AAR"

var (
	newRelic *newrelic.Application
)

// SetupNewRelic sets up monitoring in NewRelic
func SetupNewRelic(licenseKey string) {
	config := newrelic.NewConfig(appName, licenseKey)

	var err error
	newRelicApp, err := newrelic.NewApplication(config)
	newRelic = &newRelicApp

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting New Relic: %q", err)
		os.Exit(1)
	}
}
