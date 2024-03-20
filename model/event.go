package model

import (
	"fmt"
	"os"
	"time"
)


type Event struct {
	id int
	name string
	description string
	active bool
	updated string
}


func eventInit() {

	dbCreateTable("event", string(`
		CREATE TABLE event (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			active BOOL,
			updated TEXT
		)
	`))

}


func EventCreate(name string, desc string) {

	_, err := db.Query(string(`
		INSERT INTO event
			(name,description,active,updated)
		VALUES
			(?,?,?,?)
	`), name, desc, true, time.Now().Format(time.DateTime))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}

}