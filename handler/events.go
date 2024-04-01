package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/wolv89/gotnsapp/middleware"
	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
)

func GetEvents(w http.ResponseWriter, req *http.Request) {

	var eventList []model.Event
	var err error

	if middleware.IsAdmin(req) {
		eventList, err = model.GetAllEvents()
	} else {
		eventList, err = model.GetActiveEvents()
	}

	if err != nil {
		if err.Error() == "none" {
			util.HttpSuccess(w, "")
			return
		} else {
			util.HttpServerError(w, err.Error())
			return
		}
	}

	output, jerr := json.Marshal(eventList)

	if jerr != nil {
		util.HttpServerError(w, jerr.Error())
		return
	}

	util.HttpSuccess(w, string(output))

}

func CreateEvent(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	type NewEvent struct {
		Name 	string 	`json:"name"`
		Desc	string 	`json:"desc"`
		State 	bool 	`json:"state"`
	}

	var newEvent NewEvent

	err = json.Unmarshal(data, &newEvent)

	if err != nil {
		util.HttpBadRequest(w, "Parse error")
		return
	}

	result, response := model.EventCreate(newEvent.Name, newEvent.Desc, newEvent.State)

	if result {
		util.HttpJ(w, result, "", response, nil)
	} else {
		util.HttpJ(w, result, response, "", nil)
	}

}


func GetEvent(w http.ResponseWriter, req *http.Request) {

	path := req.PathValue("eventname")

	if len(path) <= 0 {
		util.HttpBadRequest(w, "No path provided")
		return
	}

	event, err := model.GetEventByPath(path)

	if err != nil {
		util.HttpJ(w, false, err.Error(), "", nil)
		return
	}

	if event.Active != true {
		if middleware.IsAdmin(req) != true {
			util.HttpJ(w, false, "Event not found", "", nil)
			return
		}
	}

	util.HttpJ(w, true, "", "", event)

}
