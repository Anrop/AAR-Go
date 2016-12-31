package main

import (
	"aar"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	maxDatabaseConnections, _ := strconv.Atoi(os.Getenv("MAX_DATABASE_CONNECTIONS"))
	newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")

	if port == "" {
		port = "8080"
	}

	var err error
	pgConfig, err := pgx.ParseURI(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid database url: %q", err)
		os.Exit(1)
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConfig,
		MaxConnections: maxDatabaseConnections,
	}
	aar.DB, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening database: %q", err)
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/missions", aar.MissionsHandler)
	r.HandleFunc("/missions/{missionId}", aar.MissionHandler)
	r.HandleFunc("/missions/{missionId}/events", aar.EventsHandler)
	r.HandleFunc("/missions/{missionId}/players", aar.MissionPlayersHandler)
	r.HandleFunc("/players", aar.PlayersHandler)

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	if newRelicLicenseKey != "" {
		aar.SetupNewRelic(newRelicLicenseKey)
	}

	// Bind to a port and pass our router in
	http.ListenAndServe(":"+port, handler)
}
