package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/wolv89/gotnsapp/middleware"
	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
	"github.com/wolv89/gotnsapp/view"
)

func GetDivisions(w http.ResponseWriter, req *http.Request) {

	event, err := strconv.Atoi(req.PathValue("eventid"))

	if err != nil || event <= 0 {
		util.HttpBadRequest(w, "No valid event provided")
		return
	}

	var divList []model.Division

	if middleware.IsAdmin(req) {
		divList, err = model.GetAllDivisions(event)
	} else {
		divList, err = model.GetActiveDivisions(event)
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

	output, jerr := json.Marshal(divList)

	if jerr != nil {
		util.HttpServerError(w, jerr.Error())
		return
	}

	util.HttpSuccess(w, string(output))

}



func CreateDivision(w http.ResponseWriter, req *http.Request) {

	event, err := strconv.Atoi(req.PathValue("eventid"))

	if err != nil || event <= 0 {
		util.HttpBadRequest(w, "No valid event provided")
		return
	}

	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	type NewDivision struct {
		Name 	string 	`json:"name"`
		State 	bool 	`json:"state"`
		Style 	int 	`json:"style"`
		Teams 	bool 	`json:"teams"`
	}

	var newDivision NewDivision

	err = json.Unmarshal(data, &newDivision)

	if err != nil {
		util.HttpBadRequest(w, "Parse error")
		return
	}

	result, response := model.DivisionCreate(event, newDivision.Name, newDivision.State, newDivision.Style, newDivision.Teams)

	if result {
		util.HttpJ(w, result, "", response, nil)
	} else {
		util.HttpJ(w, result, response, "", nil)
	}

}


func GetDivision(w http.ResponseWriter, req *http.Request) {

	eid, err := strconv.Atoi(req.PathValue("eventid"))
	dpath := req.PathValue("divisionname")

	if err != nil || eid <= 0 || len(dpath) <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	if model.GetEventStatus(eid) != true {
		if middleware.IsAdmin(req) != true {
			util.HttpJ(w, false, "Event not found", "", nil)
			return
		}
	}

	var division model.Division
	division, err = model.GetDivisionByPath(eid, dpath)

	if err != nil {
		util.HttpJ(w, false, err.Error(), "", nil)
		return
	}

	if division.Active != true {
		if middleware.IsAdmin(req) != true {
			util.HttpJ(w, false, "Division not found", "", nil)
			return
		}
	}

	util.HttpJ(w, true, "", "", division)

}


func GetDivisionStatus(w http.ResponseWriter, req *http.Request) {

	div, err := strconv.Atoi(req.PathValue("divisionid"))

	if err != nil || div <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	status := model.DivisionStatus{Entrants: false, Matches: false}

	entrantCount := model.CountEntrants(div)
	matchCount := model.CountMatches(div)

	// Handle -1 errors ?

	if entrantCount > 0 {
		status.Entrants = true
	}

	if matchCount > 0 {
		status.Matches = true
	}

	util.HttpJ(w, true, "", "", status)

}


func UpdateDivision(w http.ResponseWriter, req *http.Request) {

	div, err := strconv.Atoi(req.PathValue("divisionid"))

	if err != nil || div <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	defer req.Body.Close()
	var data []byte
	data, err = io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	type DivisionUpdate struct {
		Field 	string 	`json:"field"`
		Value 	string 	`json:"value"`
	}

	var divisionUpdates []DivisionUpdate

	err = json.Unmarshal(data, &divisionUpdates)

	if err != nil {
		util.HttpServerError(w, "Parse error")
		return
	}

	var uerr error

	for _, update := range divisionUpdates {
		uerr = model.UpdateDivision(div, update.Field, update.Value)
		if uerr != nil {
			util.HttpServerError(w, "Unable to apply update: " + update.Field)
			return
		}
	}

	util.HttpSuccess(w, "Woohoo")

}


func GetDivisionView(w http.ResponseWriter, req *http.Request) {

	div, err := strconv.Atoi(req.PathValue("divisionid"))

	if err != nil || div <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	var matches []model.Match

	view := view.MatchesView(matches)
	view.Render(context.Background(), w)

}