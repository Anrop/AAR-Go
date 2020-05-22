package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Anrop/AAR-Go/internal/aar"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	newRelicLicenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")

	if port == "" {
		port = "8080"
	}

	db, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening database: %q", err)
		os.Exit(1)
	}
	aar.DB = db

	r := mux.NewRouter()

	if newRelicLicenseKey != "" {
		aar.SetupNewRelic(newRelicLicenseKey, r)
	}

	r.HandleFunc("/missions", aar.MissionsHandler)
	r.HandleFunc("/missions/{missionId}", aar.MissionHandler)
	r.HandleFunc("/missions/{missionId}/events", aar.EventsHandler)
	r.HandleFunc("/missions/{missionId}/players", aar.MissionPlayersHandler)
	r.HandleFunc("/players", aar.PlayersHandler)
	r.HandleFunc("/players/{playerId}/missions", aar.PlayerMissionsHandler)

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	// Bind to a port and pass our router in
	http.ListenAndServe(":"+port, handler)
}
