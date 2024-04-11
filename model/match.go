package model


import (
	"fmt"
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
