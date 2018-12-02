package aar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func outputEvents(missionID int, limit int, offset int, w http.ResponseWriter) error {
	// nil value for limitStr queries all events
	var limitStr *string
	if limit > 0 {
		limitVal := strconv.Itoa(limit)
		limitStr = &limitVal
	}

	rows, err := DB.Query(`
		SELECT
			id,
			data,
			timestamp,
			type
		FROM events
		WHERE mission_id = $1
		ORDER BY timestamp ASC
		LIMIT $2
		OFFSET $3
	`, missionID, limitStr, offset)

	if err != nil {
		return err
	}
	defer rows.Close()

	enc := json.NewEncoder(w)
	w.Write([]byte("["))

	var first = true

	for rows.Next() {
		if first {
			first = false
		} else {
			w.Write([]byte(","))
		}

		event := Event{}
		err := rows.Scan(&event.ID, &event.Data, &event.Timestamp, &event.Type)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading event row from database: %v", err)
			continue
		}

		// Remove Rails prefix from Type
		event.Type = strings.Replace(event.Type, "Events::", "", 1)

		// Move properties inline to event object
		event.Player = event.Data.Player
		event.Projectile = event.Data.Projectile
		event.Unit = event.Data.Unit
		event.Vehicle = event.Data.Vehicle
		event.Data = nil

		enc.Encode(event)
	}

	w.Write([]byte("]"))

	return nil
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	if newRelic != nil {
		txn := (*newRelic).StartTransaction(r.URL.Path, w, r)
		defer txn.End()
	}

	params := mux.Vars(r)
	missionID, err := strconv.Atoi(params["missionId"])
	if err != nil {
		http.Error(w, "Invalid mission id", http.StatusBadRequest)
		fmt.Fprintf(os.Stderr, "Error parsing mission id: %v", err)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if err := outputEvents(missionID, limit, offset, w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error reading events: %v", err)
	}
}
