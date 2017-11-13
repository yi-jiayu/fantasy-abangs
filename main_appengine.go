// +build appengine

package main

import "net/http"

func init() {
	http.HandleFunc("/api/players", getPlayers)
}
