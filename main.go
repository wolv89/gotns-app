package main

import (
	"fmt"
	"net/http"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/wolv89/gotnsapp/model"

)


func main() {

	fmt.Println("## Server Init")

	model.DatabaseInit()
	defer model.DatabaseClose()

	fmt.Println("## Database connected, launching...")


	fmt.Println("## Server Started")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "GOTNS")
	})

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

	http.ListenAndServe(":8040", nil)

}
