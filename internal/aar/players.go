package aar

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func outputPlayersFromRows(rows pgx.Rows, w http.ResponseWriter) error {
	enc := json.NewEncoder(w)
	w.Write([]byte("["))

	var first = true

	for rows.Next() {
		if first {
			first = false
		} else {
			w.Write([]byte(","))
		}

		player := Player{}
		err := rows.Scan(&player.Name, &player.UID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading player row from database: %v", err)
			continue
		}

		enc.Encode(player)
	}

	w.Write([]byte("]"))

	return nil
}

func outputPlayers(ctx context.Context, w http.ResponseWriter) error {
	rows, err := DB.Query(ctx, `
		SELECT
			data #>> '{player, name}' as name,
			data #>> '{player, uid}' as uid
		FROM events
		WHERE
			data #>> '{player, name}' IS NOT NULL
			AND
			data #>> '{player, uid}' IS NOT NULL
		GROUP BY uid, name
		ORDER BY name
	`)

	if err != nil {
		return err
	}
	defer rows.Close()

	return outputPlayersFromRows(rows, w)
}

func outputMissionPlayers(ctx context.Context, missionID int, w http.ResponseWriter) error {
	rows, err := DB.Query(ctx, `
		SELECT
			data #>> '{player, name}' as name,
			data #>> '{player, uid}' as uid
		FROM events
		WHERE
			mission_id = $1
			AND
			data #>> '{player, name}' IS NOT NULL
			AND
			data #>> '{player, uid}' IS NOT NULL
		GROUP BY uid, name
		ORDER BY name
	`, missionID)

	if err != nil {
		return err
	}
	defer rows.Close()

	return outputPlayersFromRows(rows, w)
}

// PlayersHandler is used to handle the players endpoint
func PlayersHandler(w http.ResponseWriter, r *http.Request) {
	if err := outputPlayers(r.Context(), w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error reading missions: %v", err)
	}
}

// MissionPlayersHandler is used to handle the mission players endpoint
func MissionPlayersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	missionID, err := strconv.Atoi(params["missionId"])
	if err != nil {
		http.Error(w, "Invalid mission id", http.StatusBadRequest)
		fmt.Fprintf(os.Stderr, "Error parsing mission id: %v", err)
		return
	}

	if err := outputMissionPlayers(r.Context(), missionID, w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error reading players: %v", err)
	}
}
