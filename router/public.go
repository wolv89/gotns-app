package router

import (
	"net/http"

	"github.com/wolv89/gotnsapp/handler"
	"github.com/wolv89/gotnsapp/util"
)

func LoadPublicRoutes(router *http.ServeMux) {

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		util.HttpBadRequest(w, "")
	})

	router.HandleFunc("/login", handler.Login)

	router.HandleFunc("/events", handler.GetEvents)
	router.HandleFunc("/event/{eventname}", handler.GetEvent)

	router.HandleFunc("/divisions/{eventid}", handler.GetDivisions)
	router.HandleFunc("/event/{eventid}/division/{divisionname}", handler.GetDivision)

	router.HandleFunc("/division/{divisionid}/view", handler.GetDivisionView)

	// router.HandleFunc("/toolbox/dbnuke", handler.ToolboxNuke)
	// router.HandleFunc("/toolbox/dblaunch", handler.ToolboxLaunch)

}