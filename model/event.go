package model

import (
	"fmt"
	"os"

	"github.com/wolv89/gotnsapp/util"
)


type Event struct {
	id int
	name string
	path string
	description string
	active bool
	updated string
}


func eventInit() {

	dbCreateTable("event", `
		CREATE TABLE event (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			path TEXT,
			description TEXT,
			active BOOL,
			updated TEXT
		)
	`)

}


func EventCreate(name string, desc string) (bool, string) {

	if name == "" {
		return false, "Bad request"
	}

	path := util.StringToPath(name)
	query := db.QueryRow(fmt.Sprintf(`SELECT id FROM event WHERE name = "%s" OR path = "%s"`, name, path))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return false, query.Err().Error()
	}

	var id int
	query.Scan(&id)

	if id > 0 {
		return false, "An event with that name or path already exists"
	}

	_, err := db.Query(`
		INSERT INTO event
			(name,path,description,active,updated)
		VALUES 
			(?,?,?,?,?)`,
		name, path, desc, true, dbNow())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		return false, err.Error()
	}

	return true, ""

}