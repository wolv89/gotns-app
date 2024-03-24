package router

import (
	"net/http"

	"github.com/wolv89/gotnsapp/handler"
)

func LoadAdminRoutes(router *http.ServeMux) {

	router.HandleFunc("GET /logout", handler.Logout)

}
