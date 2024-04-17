package router

import (
	"net/http"

	"github.com/wolv89/gotnsapp/handler"
)

func LoadAdminRoutes(router *http.ServeMux) {

	router.HandleFunc("/ready", handler.Ready)
	router.HandleFunc("/logout", handler.Logout)

	router.HandleFunc("POST /event/new", handler.CreateEvent)

	router.HandleFunc("POST /event/{eventid}/division/new", handler.CreateDivision)

	router.HandleFunc("/players", handler.GetPlayers)
	router.HandleFunc("POST /player/new", handler.CreatePlayer)

	router.HandleFunc("/division/{divisionid}/status", handler.GetDivisionStatus)
	router.HandleFunc("/division/{divisionid}/entrants", handler.GetEntrants)
	router.HandleFunc("POST /division/{divisionid}/entrants/new", handler.CreateEntrants)
	router.HandleFunc("POST /division/{divisionid}/matches/new", handler.CreateMatches)

}
