package aar

import "time"

// Mission data structure
type Mission struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Length    *int32    `json:"length"`
	Name      string    `json:"name"`
	World     string    `json:"world"`
}

// Event data structure
type Event struct {
	ID         int32       `json:"id"`
	Type       string      `json:"type"`
	Data       *EventData  `json:"data,omitempty"`
	Timestamp  time.Time   `json:"timestamp"`
	Player     *Player     `json:"player,omitempty"`
	Projectile *Projectile `json:"projectile,omitempty"`
	Unit       *Unit       `json:"unit,omitempty"`
	Vehicle    *Vehicle    `json:"vehicle,omitempty"`
}

// EventData data structure
type EventData struct {
	Player     *Player     `json:"player,omitempty"`
	Projectile *Projectile `json:"projectile,omitempty"`
	Unit       *Unit       `json:"unit,omitempty"`
	Vehicle    *Vehicle    `json:"vehicle,omitempty"`
}

// Player data structure
type Player struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}

// Position data structure
type Position struct {
	Dir float64 `json:"dir"`
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
	Z   float64 `json:"z"`
}

// Projectile data structure
type Projectile struct {
	ID         string   `json:"id"`
	Position   Position `json:"position"`
	Side       string   `json:"side"`
	Simulation string   `json:"simulation"`
}

// Unit data structure
type Unit struct {
	ID        string   `json:"id"`
	LifeState string   `json:"life_state"`
	Name      string   `json:"name"`
	Position  Position `json:"position"`
	Side      string   `json:"side"`
	VehicleID string   `json:"vehicle_id"`
}

// Vehicle data structure
type Vehicle struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Position   Position `json:"position"`
	Side       string   `json:"side"`
	Simulation string   `json:"simulation"`
}
