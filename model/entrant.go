package model

import (
	"errors"
	"fmt"
	"strings"
)


type Entrant struct {
	Id int			`json:"id"`
	Division int	`json:"division"`
	Player1 int		`json:"player1"`
	Player2 int		`json:"player2"`
	Team bool		`json:"team"`
	Seed int		`json:"seed"`
}


type NamedEntrant struct {
	Id int
	P1Id int
	P1FirstName string
	P1LastName string
	P2Id int
	P2FirstName string
	P2LastName string
	Team bool
	Seed int
}


func entrantInit() {

	dbCreateTable("entrant", string(`
		CREATE TABLE entrant (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			division INTEGER,
			player1 INTEGER,
			player2 INTEGER,
			team BOOL,
			seed INTEGER
		)
	`))

}


func GetSingleEntrant(div int, p1 int) (Entrant, error) {

	var entrant Entrant

	if div <= 0 || p1 <= 0 {
		return entrant, errors.New("Bad request")
	}

	query := db.QueryRow(fmt.Sprintf(`
		SELECT * FROM entrant
		WHERE division = %d AND player1 = %d
		`, div, p1))

	if query.Err() != nil {
		return entrant, query.Err()
	}

	query.Scan(&entrant.Id, &entrant.Division, &entrant.Player1, &entrant.Player2, &entrant.Team, &entrant.Seed)

	if entrant.Id <= 0 {
		return entrant, errors.New("Not found")
	}

	return entrant, nil

}


func GetTeamEntrant(div int, p1 int, p2 int) (Entrant, error) {

	var entrant Entrant

	if div <= 0 || p1 <= 0 || p2 <= 0 {
		return entrant, errors.New("Bad request")
	}

	query := db.QueryRow(fmt.Sprintf(`
		SELECT * FROM entrant
		WHERE division = %d AND player1 = %d AND player2 = %d
		`, div, p1, p2))

	if query.Err() != nil {
		return entrant, query.Err()
	}

	query.Scan(&entrant.Id, &entrant.Division, &entrant.Player1, &entrant.Player2, &entrant.Team, &entrant.Seed)

	if entrant.Id <= 0 {
		return entrant, errors.New("Not found")
	}

	return entrant, nil

}




func CreateSingleEntrant(div int, p1 int, seed int) (bool, int) {

	if p1 <= 0 || div <= 0 {
		return false, 0
	}

	entrant, err := GetSingleEntrant(div,p1)

	_, err = db.Query(`
		INSERT INTO entrant
			(division,player1,player2,team,seed)
		VALUES
			(?,?,?,?,?)`,
			div, p1, 0, false, seed)

	if err != nil {
		return false, 0
	}

	entrant, err = GetSingleEntrant(div,p1)

	return true, entrant.Id

}



func CreateTeamEntrant(div int, p1 int, p2 int, seed int) (bool, int) {

	if p1 <= 0 || p2 <= 0 || div <= 0 {
		return false, 0
	}

	entrant, err := GetTeamEntrant(div,p1,p2)

	_, err = db.Query(`
		INSERT INTO entrant
			(division,player1,player2,team,seed)
		VALUES
			(?,?,?,?,?)`,
			div, p1, p2, true, seed)

	if err != nil {
		return false, 0
	}

	entrant, err = GetTeamEntrant(div,p1,p2)

	return true, entrant.Id

}



func CountEntrants(div int) int {

	if div <= 0 {
		return -1
	}

	sizequery := db.QueryRow(fmt.Sprintf(`
		SELECT COUNT(id) FROM entrant
		WHERE division = %d
		`, div))

	if sizequery.Err() != nil {
		return -1
	}

	var entrantCount int
	sizequery.Scan(&entrantCount)

	return entrantCount

}




func GetEntrants(div int) ([]Entrant, error) {

	var list []Entrant

	if div <= 0 {
		return list, errors.New("Bad request")
	}

	entrantCount := CountEntrants(div)

	if entrantCount == 0 {
		return list, errors.New("None")
	} else if entrantCount < 0 {
		return list, errors.New("Query error")
	}

	listquery, err := db.Query(fmt.Sprintf(`
		SELECT *
		FROM entrant
		WHERE division = %d
		ORDER BY id ASC
	`, div))

	defer listquery.Close()

	if err != nil {
		return list, err
	}

	list = make([]Entrant, 0, entrantCount)

	for listquery.Next() {

		var entrant Entrant

		if qerr := listquery.Scan(&entrant.Id, &entrant.Division, &entrant.Player1, &entrant.Player2, &entrant.Team, &entrant.Seed); qerr != nil {
			return list, qerr
		}

		list = append(list, entrant)

	}

	return list, nil

}



func GetEntrantName(id int) string {

	if id <= 0 {
		return "---"
	}

	var entrant NamedEntrant

	query := db.QueryRow(fmt.Sprintf(`
		SELECT e.id, e.team, e.seed, p1.id, p1.firstname, p1.lastname, p2.id, p2.firstname, p2.lastname
		FROM entrant AS e
		LEFT JOIN player AS p1
			ON p1.id = e.player1
		LEFT JOIN player AS p2
			ON p2.id = e.player2
		WHERE e.id = %d
		`, id))

	if query.Err() != nil {
		return "Unknown player(s)"
	}

	query.Scan(&entrant.Id, &entrant.Team, &entrant.Seed, &entrant.P1Id, &entrant.P1FirstName, &entrant.P1LastName, &entrant.P2Id, &entrant.P2FirstName, &entrant.P2LastName)

	if entrant.Id != id {
		return "Unknown player(s)"
	}

	var name strings.Builder

	if entrant.Seed > 0 {
		name.WriteString(fmt.Sprintf("<span class=\"seed\">%d</span>", entrant.Seed))
	}

	name.WriteString(fmt.Sprintf("<span class=\"player\">%s. %s</span>", entrant.P1FirstName[:1], entrant.P1LastName))

	if entrant.Team {
		name.WriteString(fmt.Sprintf("<span class=\"player\">%s. %s</span>", entrant.P2FirstName[:1], entrant.P2LastName))
	}

	return name.String()

}
