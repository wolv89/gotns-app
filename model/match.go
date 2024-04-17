package model

import (
	"errors"
	"fmt"
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
