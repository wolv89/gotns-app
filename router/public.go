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


	/*

	http.HandleFunc("/launch", func(w http.ResponseWriter, req *http.Request){
		model.DatabaseLaunch()
		fmt.Fprintf(w, "Database Launch")
	})

	http.HandleFunc("/nuke", func(w http.ResponseWriter, req *http.Request){
		model.DatabaseNuke()
		fmt.Fprintf(w, "Database Nuke")
	})

	http.HandleFunc("/event/create", func(w http.ResponseWriter, req *http.Request){
		q := req.URL.Query()
		eventName := q.Get("name")

		res, err := model.EventCreate(eventName, "Clay boy")
		if !res {
			fmt.Fprintf(w, fmt.Sprintf("Unable to create event: [%s] {%s}", eventName, err))
		} else {
			fmt.Fprintf(w, fmt.Sprintf("Create event: [%s]", eventName))
		}
	})

	http.HandleFunc("/event/list", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "List events...")
	})

	http.HandleFunc("/get-word", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "<strong>Word</strong>")
	})

	*/

}