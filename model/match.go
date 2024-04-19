package model

import (
	"errors"
	"fmt"
	"math"
)


type MatchStatus int

const (
	MatchBlank MatchStatus = iota
	MatchReady
	MatchPlaying
	MatchFinished
)



type Match struct {
	Id int			`json:"id"`
	Division int	`json:"division"`
	Entrant1 int	`json:"entrant1"`
	Entrant2 int	`json:"entrant2"`
	Score string	`json:"score"`
	Notes string	`json:"notes"`
	Seq int			`json:"seq"`
	Start string	`json:"start"`
	Updated string	`json:"updated"`
	Status int		`json:"status"`
	Winner int		`json:"winner"`
}


func matchInit() {

	dbCreateTable("match", string(`
		CREATE TABLE match (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			division INTEGER,
			entrant1 INTEGER,
			entrant2 INTEGER,
			score TEXT,
			notes TEXT,
			seq INTEGER,
			start TEXT,
			updated TEXT,
			status INTEGER,
			winner INTEGER
		)
	`))

}


func MatchCreate(div int, e1 int, e2 int, seq int, status MatchStatus) error {

	if div <= 0 {
		return errors.New("Bad request")
	}

	_, err := db.Query(`
		INSERT INTO match
			(division,entrant1,entrant2,score,notes,seq,start,updated,status,winner)
		VALUES 
			(?,?,?,?,?,?,?,?,?,?)`,
		div, e1, e2, "", "", seq, "", dbNow(), status, 0)

	if err != nil {
		return err
	}

	return nil

}


func MatchUpdate(mid int, score string, notes string, start string, status int, winner int) error {

	if mid <= 0 {
		return errors.New("Bad request")
	}

	query, err := db.Query(fmt.Sprintf(`
		UPDATE match
		SET score = "%s", notes = "%s", start = "%s", status = %d, winner = %d
		WHERE id = %d
	`, score, notes, start, status, winner, mid))

	defer query.Close()

	if err != nil {
		return err
	}

	return nil

}



func GetMatch(mid int) (Match, error) {

	var match Match

	if mid <= 0 {
		return match, errors.New("No ID provided")
	}

	query := db.QueryRow(fmt.Sprintf(`SELECT * FROM match WHERE id = %d`, mid))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return match, query.Err()
	}

	if qerr := query.Scan(&match.Id, &match.Division, &match.Entrant1, &match.Entrant2, &match.Score, &match.Notes, &match.Seq, &match.Start, &match.Updated, &match.Status, &match.Winner); qerr != nil {
		return match, qerr
	}

	return match, nil

}


func GetNextMatch(div int, seq int) (Match, error) {

	var match Match

	if div <= 0 || seq < 0 {
		return match, errors.New("Bad search")
	}

	query := db.QueryRow(fmt.Sprintf(`SELECT * FROM match WHERE division = %d AND seq = %d`, div, seq))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return match, query.Err()
	}

	if qerr := query.Scan(&match.Id, &match.Division, &match.Entrant1, &match.Entrant2, &match.Score, &match.Notes, &match.Seq, &match.Start, &match.Updated, &match.Status, &match.Winner); qerr != nil {
		return match, qerr
	}

	return match, nil

}


func (m *Match) SetEntrant(which int, entrant int) {

	if entrant <= 0 {
		return
	}

	target := "entrant1"
	if which == 2 {
		target = "entrant2"
	}

	query, _ := db.Query(fmt.Sprintf(`
		UPDATE match
		SET %s = %d
		WHERE id = %d
	`, target, entrant, m.Id))

	defer query.Close()

}


func (m *Match) SetStatus(status MatchStatus) {

	query, _ := db.Query(fmt.Sprintf(`
		UPDATE match
		SET status = %d
		WHERE id = %d
	`, int(status), m.Id))

	defer query.Close()

}




func CountMatches(div int) int {

	if div <= 0 {
		return -1
	}

	sizequery := db.QueryRow(fmt.Sprintf(`
		SELECT COUNT(id) FROM match
		WHERE division = %d
		`, div))

	if sizequery.Err() != nil {
		return -1
	}

	var matchCount int
	sizequery.Scan(&matchCount)

	return matchCount

}




func GetMatches(div int) ([]Match, error) {

	var matches []Match

	if div <= 0 {
		return matches, errors.New("Invalid division ID")
	}

	query, err := db.Query(fmt.Sprintf(`
		SELECT *
		FROM match
		WHERE division = %d
		ORDER BY seq ASC
	`, div))

	defer query.Close()

	if err != nil {
		return matches, err
	}

	for query.Next() {

		var match Match

		if err = query.Scan(&match.Id,&match.Division,&match.Entrant1,&match.Entrant2,&match.Score,&match.Notes,&match.Seq,&match.Start,&match.Updated,&match.Status,&match.Winner); err != nil {
			return matches, err
		}

		matches = append(matches, match)

	}

	return matches, nil

}




func (m Match) GetRoundName() string {

	c := int(math.Log2(float64(m.Seq + 1)))

	switch c {
		case 0:
			return "Final"
		case 1:
			return "Semi Final"
		case 2:
			return "Quarter Final"
		case 3:
			return "Round of 8"
		case 4:
			return "Round of 16"
		case 5:
			return "Round of 32"
		case 6:
			return "Round of 64"
		case 7:
			return "Round of 128"
	}

	return "That's unpossible"

}


func (m Match) GetStatus() string {

	switch m.Status {
		case int(MatchBlank):
			return ""
		case int(MatchReady):
			return "Upcoming"
		case int(MatchPlaying):
			return "Playing"
		case int(MatchFinished):
			return "Complete"
	}

	return ""

}


func (m Match) GetStatusSlug() string {

	switch m.Status {
		case int(MatchBlank):
			return "status-na"
		case int(MatchReady):
			return "status-ready"
		case int(MatchPlaying):
			return "status-live"
		case int(MatchFinished):
			return "status-done"
	}

	return ""

}