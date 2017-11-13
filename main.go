// +build !appengine

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/players", getPlayers)

	panic(http.ListenAndServe(":8080", nil))
}
