package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
	"github.com/wolv89/gotnsapp/view"
)


func CreateMatches(w http.ResponseWriter, req *http.Request) {

	division, err := strconv.Atoi(req.PathValue("divisionid"))

	if division <= 0 || err != nil {
		util.HttpBadRequest(w, "No valid division provided")
		return
	}

	defer req.Body.Close()
	var data []byte
	data, err = io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	type NewMatch struct {
		Entrant1		int 	`json:"entrant1"`
		Entrant2		int 	`json:"entrant2"`
	}

	var newMatches []NewMatch

	err = json.Unmarshal(data, &newMatches)

	if err != nil {
		util.HttpServerError(w, "Parse error")
		return
	}

	structure := len(newMatches) - 1
	var seq int

	// Loading the match tree from the final down, all blank to start
	for seq = 0; seq < structure; seq++ {

		err = model.MatchCreate(division, 0, 0, seq, model.MatchBlank)

		if err != nil {
			util.HttpServerError(w, "Could not create match")
			return
		}

	}

	// Load round 1 matches
	for _, match := range newMatches {

		err = model.MatchCreate(division, match.Entrant1, match.Entrant2, seq, model.MatchReady)
		seq++

		if err != nil {
			util.HttpServerError(w, "Could not create match")
			return
		}

	}

	util.HttpSuccess(w, "Done")

}





func GetMatchesView(w http.ResponseWriter, req *http.Request) {

	div, err := strconv.Atoi(req.PathValue("divisionid"))

	if err != nil || div <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	var matches []model.Match
	matches, err = model.GetMatches(div)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	mcount := len(matches)							// Match count

	if mcount <= 0 {
		util.HttpBadRequest(w, "No matches")
		return
	}

	dcount := (mcount + 1) / 2						// Draw count
	rcount := int(math.Log2(float64(dcount))) + 1	// Rounds count

	rounds := make([][]int, rcount)
	rm := dcount
	ri, min, max := 0, 0, 0

	for r := 0; r < rcount; r++ {

		rounds[r] = make([]int, rm)
		min = rm - 1
		max = min + rm

		for ri = min; ri < max; ri++ {
			rounds[r][ri - min] = ri
		}

		rm /= 2

	}

	view := view.MatchesView(matches, rounds)
	view.Render(context.Background(), w)

}




func GetMatch(w http.ResponseWriter, req *http.Request) {

	mid, err := strconv.Atoi(req.PathValue("matchid"))

	if err != nil || mid <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	var match model.Match
	match, err = model.GetMatch(mid)

	if err != nil {
		util.HttpBadRequest(w, err.Error())
		return
	}

	type MatchEdit struct {
		Match model.Match				`json:"match"`
		Entrant1 model.NamedEntrant		`json:"entrant1"`
		Entrant2 model.NamedEntrant		`json:"entrant2"`
	}

	e1, _ := model.GetNamedEntrant(match.Entrant1)
	e2, _ := model.GetNamedEntrant(match.Entrant2)

	edit := MatchEdit{Match: match, Entrant1: e1, Entrant2: e2}

	output, jerr := json.Marshal(edit)

	if jerr != nil {
		util.HttpServerError(w, jerr.Error())
		return
	}

	util.HttpSuccess(w, string(output))

}




func UpdateMatch(w http.ResponseWriter, req *http.Request) {

	mid, err := strconv.Atoi(req.PathValue("matchid"))

	if err != nil || mid <= 0 {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)

	if err != nil {
		util.HttpBadRequest(w, "Bad request")
		return
	}

	var match model.Match
	err = json.Unmarshal(data, &match)

	if err != nil {
		util.HttpBadRequest(w, "Parse issue")
		return
	}

	err = model.MatchUpdate(match.Id, match.Score, match.Notes, match.Start, match.Status, match.Winner)

	if err != nil {
		util.HttpServerError(w, err.Error())
		return
	}

	err = progressEntrant(match)

	if err != nil {
		util.HttpServerError(w, err.Error())
		return
	}

	util.HttpSuccess(w, "Saved")

}


func progressEntrant(match model.Match) error {

	if match.Winner <= 0 || match.Seq <= 0 {
		return nil
	}

	even := match.Seq % 2 == 0
	next := 0

	if even {
		next = (match.Seq - 2)/2
	} else {
		next = (match.Seq - 1)/2
	}

	nextMatch, err := model.GetNextMatch(match.Division, next)

	if err != nil {
		return err
	}

	var entrant int
	if match.Winner == 1 {
		entrant = match.Entrant1
	} else if match.Winner == 2 {
		entrant = match.Entrant2
	}

	if entrant <= 0 {
		return errors.New("No winner")
	}

	if even {
		if nextMatch.Entrant2 == 0 {
			nextMatch.SetEntrant(2, entrant)
			if nextMatch.Entrant1 != 0 {
				nextMatch.SetStatus(model.MatchReady)
			}
		}
	} else {
		if nextMatch.Entrant1 == 0 {
			nextMatch.SetEntrant(1, entrant)
			if nextMatch.Entrant2 != 0 {
				nextMatch.SetStatus(model.MatchReady)
			}
		}
	}

	return nil

}