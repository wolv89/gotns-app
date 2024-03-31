package router

import (
	"net/http"

	"github.com/wolv89/gotnsapp/handler"
)

func LoadAdminRoutes(router *http.ServeMux) {

	router.HandleFunc("/ready", handler.Ready)
	router.HandleFunc("/logout", handler.Logout)

	router.HandleFunc("POST /event/new", handler.CreateEvent)

}
