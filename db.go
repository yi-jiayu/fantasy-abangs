package main

import (
	"encoding/json"
	"net/http"
)

type Player struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cost      int    `json:"cost"`
	Points    int    `json:"points"`
}

var players = []Player{
	{
		FirstName: "Edmond",
		LastName:  "To",
		Cost:      60,
		Points:    0,
	},
	{
		FirstName: "Andrew",
		LastName:  "Sng",
		Cost:      60,
		Points:    0,
	},
	{
		FirstName: "Tu Yong",
		LastName:  "Tan",
		Cost:      60,
		Points:    0,
	},
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(players)
	if err != nil {
		panic(err)
	}
}
