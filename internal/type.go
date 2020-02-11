package internal

import (
	"time"
)

var AcolyteNames = map[string]int{
	"Angst":    1,
	"Malice":   2,
	"Mania":    3,
	"Misery":   4,
	"Torment":  5,
	"Violence": 6,
}

// Acolyte ...
type Acolyte struct {
	Name                   string
	Discovered             bool
	AgentType              string
	HealthPercent          float64
	LastDiscoveredTime     int64
	LastDiscoveredLocation Node
	Mods                   map[string]float64
	notified               bool
}

// Node
type Node struct {
	Name        string `json:"value"`
	Enemy       string `json:"enemy"`
	MissionType string `json:"type"`
}

// Tracker ...
type Tracker struct {
	cwd      string
	silent   bool
	notify   bool
	timer    *Timer
	acolytes map[string]*Acolyte
	nodes    map[string]Node
}

// Timer
type Timer struct {
	status   string
	duration time.Duration
	end      time.Time
	left     time.Duration
}
