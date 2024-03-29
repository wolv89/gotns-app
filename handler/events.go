package handler

import (
	"net/http"
	"encoding/json"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
)

func GetEvents(w http.ResponseWriter, req *http.Request) {

	var eventList []model.Event
	var err error

	// If ADMIN then GetAllEvents() ?
	eventList, err = model.GetActiveEvents()

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
