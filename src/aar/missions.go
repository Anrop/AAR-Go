package aar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getMissions() ([]Mission, error) {
	rows, err := DB.Query(`
		SELECT
			id,
			created_at,
			EXTRACT(
				epoch FROM (
					SELECT timestamp
					FROM events
					WHERE events.mission_id = missions.id
					ORDER BY timestamp DESC
					LIMIT 1
				) - (
					SELECT timestamp
					FROM events
					WHERE events.mission_id = missions.id
					ORDER BY timestamp ASC
					LIMIT 1
				)
			)::int AS length,
			name,
			world
		FROM missions
		ORDER BY created_at DESC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]Mission, 0)
	for rows.Next() {
		mission := Mission{}
		e := rows.Scan(&mission.ID, &mission.CreatedAt, &mission.Length, &mission.Name, &mission.World)
		if e != nil {
			return nil, e
		}

		res = append(res, mission)
	}

	return res, nil
}

func getMission(missionId string) (*Mission, error) {
	row := DB.QueryRow("SELECT id, name, world FROM missions WHERE id = $1", missionId)
	mission := new(Mission)
	err := row.Scan(&mission.ID, &mission.Name, &mission.World)

	if err != nil {
		return nil, err
	}

	return mission, nil
}

func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	missions, err := getMissions()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(missions)
}

func MissionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	missionId := params["missionId"]

	mission, err := getMission(missionId)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(mission)
}
