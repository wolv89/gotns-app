package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
)


func GetPlayers(w http.ResponseWriter, req *http.Request) {

	players, err := model.GetPlayers()

	if err != nil {
		util.HttpServerError(w, err.Error())
		return
	}

	output, jerr := json.Marshal(players)

	if jerr != nil {
		util.HttpServerError(w, jerr.Error())
		return
	}

	util.HttpSuccess(w, string(output))

}


func CreatePlayer(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	type NewPlayer struct {
		FirstName string	`json:"firstname"`
		LastName string		`json:"lastname"`
	}

	var newPlayer NewPlayer

	err = json.Unmarshal(data, &newPlayer)

	if err != nil {
		util.HttpBadRequest(w, "Parse error")
		return
	}

	result, response := model.PlayerCreate(newPlayer.FirstName, newPlayer.LastName)

	if result {
		util.HttpJ(w, result, "", strconv.Itoa(response), nil)
	} else {
		util.HttpJ(w, result, "", "", nil)
	}

}
