package model

import (
	"errors"
	"fmt"
)


type Player struct {
	Id int				`json:"id"`
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
}


func playerInit() {

	dbCreateTable("player", string(`
		CREATE TABLE player (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			firstname TEXT,
			lastname TEXT
		)
	`))

}


func GetPlayerByName(fname string, lname string) (Player, error) {

	var player Player

	if fname == "" || lname == "" {
		return player, errors.New("Bad request")
	}

	query := db.QueryRow(fmt.Sprintf(`
		SELECT * FROM player
		WHERE firstname = "%s" AND lastname = "%s"
		`, fname, lname))

	if query.Err() != nil {
		return player, errors.New("Bad query")
	}

	query.Scan(&player.Id, &player.FirstName, &player.LastName)

	if player.Id <= 0 {
		return player, errors.New("Not found")
	}

	return player, nil

}



func PlayerCreate(fname string, lname string) (bool, int) {

	if fname == "" || lname == "" {
		return false, 0
	}

	player, err := GetPlayerByName(fname, lname)

	if err != nil {
		return false, 0
	}

	_, err = db.Query(`
		INSERT INTO player
			(firstname,lastname)
		VALUES
			(?,?)`,
		fname, lname)

	if err != nil {
		return false, 0
	}

	player, err = GetPlayerByName(fname, lname)

	return true, player.Id

}


func GetPlayers() ([]Player, error) {

	query, err := db.Query(`
		SELECT *
		FROM player
		ORDER BY lastname ASC
	`)

	defer query.Close()

	if err != nil {
		return nil, err
	}

	var list []Player

	for query.Next() {

		var player Player

		if qerr := query.Scan(&player.Id, &player.FirstName, &player.LastName); qerr != nil {
			return nil, qerr
		}

		list = append(list, player)

	}

	if err = query.Err(); err != nil {
		return nil, err
	}

	return list, nil

}
