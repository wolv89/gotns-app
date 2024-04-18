package handler

import (
	"context"
	"encoding/json"
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