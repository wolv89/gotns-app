package handler


import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
)

func GetEntrants(w http.ResponseWriter, req *http.Request) {

	division, err := strconv.Atoi(req.PathValue("divisionid"))

	if err != nil || division <= 0 {
		util.HttpBadRequest(w, "No valid division provided")
		return
	}

	entrantList, err := model.GetEntrants(division)

	if err != nil {
		if err.Error() == "none" {
			util.HttpSuccess(w, "")
			return
		} else {
			util.HttpServerError(w, err.Error())
			return
		}
	}

	output, jerr := json.Marshal(entrantList)

	if jerr != nil {
		util.HttpServerError(w, jerr.Error())
		return
	}

	util.HttpSuccess(w, string(output))

}



func CreateEntrants(w http.ResponseWriter, req *http.Request) {

	division, err := strconv.Atoi(req.PathValue("divisionid"))
	teams, terr := strconv.Atoi(req.URL.Query().Get("teams"))

	if err != nil || terr != nil || division <= 0 {
		util.HttpBadRequest(w, "No valid division provided")
		return
	}

	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	if teams == 1 {
		createTeamEntrants(w, division, data)
	} else {
		createSingleEntrants(w, division, data)
	}

}



func createSingleEntrants(w http.ResponseWriter, div int, data []byte) {

	type NewEntrant struct {
		Player1		int 	`json:"player1"`
		Seed		int 	`json:"seed"`
	}

	var newEntrants []NewEntrant

	err := json.Unmarshal(data, &newEntrants)

	if err != nil {
		util.HttpBadRequest(w, "Parse error")
		return
	}

	var added bool

	for _, entrant := range newEntrants {

		added, _ = model.CreateSingleEntrant(div, entrant.Player1, entrant.Seed)

		if !added {
			util.HttpServerError(w, "Well shit")
			return
		}

	}

	util.HttpJ(w, true, "", "", nil)

}


func createTeamEntrants(w http.ResponseWriter, div int, data []byte) {

	type NewEntrant struct {
		Player1		int 	`json:"player1"`
		Player2		int 	`json:"player2"`
		Seed		int 	`json:"seed"`
	}

	var newEntrants []NewEntrant

	err := json.Unmarshal(data, &newEntrants)

	if err != nil {
		util.HttpBadRequest(w, "Parse error")
		return
	}

	var added bool

	for _, entrant := range newEntrants {

		added, _ = model.CreateTeamEntrant(div, entrant.Player1, entrant.Player2, entrant.Seed)

		if !added {
			util.HttpServerError(w, "Well shit")
			return
		}

	}

	util.HttpJ(w, true, "", "", nil)

}