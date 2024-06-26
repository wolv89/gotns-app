package model

import (
	"errors"
	"fmt"
	"os"

	"github.com/wolv89/gotnsapp/util"
)


type Event struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Path string			`json:"path"`
	Description string	`json:"description"`
	Active bool			`json:"active"`
	Updated string		`json:"updated"`
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


func EventCreate(name string, desc string, state bool) (bool, string) {

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
		name, path, desc, state, dbNow())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		return false, err.Error()
	}

	return true, path

}


func GetActiveEvents() ([]Event, error) {

	query, err := db.Query(`
		SELECT *
		FROM event
		WHERE active = TRUE
		ORDER BY updated DESC
	`)

	defer query.Close()

	if err != nil {
		return nil, err
	}

	var list []Event

	for query.Next() {

		var event Event

		if qerr := query.Scan(&event.Id, &event.Name, &event.Path, &event.Description, &event.Active, &event.Updated); qerr != nil {
			return nil, qerr
		}

		list = append(list, event)

	}

	if len(list) <= 0 {
		return nil, errors.New("none")
	}

	if err = query.Err(); err != nil {
		return nil, err
	}

	return list, nil

}


func GetAllEvents() ([]Event, error) {

	query, err := db.Query(`
		SELECT *
		FROM event
		ORDER BY active DESC, updated DESC
	`)

	defer query.Close()

	if err != nil {
		return nil, err
	}

	var list []Event

	for query.Next() {

		var event Event

		if qerr := query.Scan(&event.Id, &event.Name, &event.Path, &event.Description, &event.Active, &event.Updated); qerr != nil {
			return nil, qerr
		}

		list = append(list, event)

	}

	if len(list) <= 0 {
		return nil, errors.New("none")
	}

	if err = query.Err(); err != nil {
		return nil, err
	}

	return list, nil

}


func GetEventByPath(path string) (Event, error) {

	var event Event

	if len(path) <= 0 {
		return event, errors.New("No path provided")
	}

	query := db.QueryRow(fmt.Sprintf(`SELECT * FROM event WHERE path = "%s"`, path))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return event, query.Err()
	}

	if qerr := query.Scan(&event.Id, &event.Name, &event.Path, &event.Description, &event.Active, &event.Updated); qerr != nil {
		return event, qerr
	}

	return event, nil

}


func GetEventStatus(id int) bool {

	if id <= 0 {
		return false
	}

	query := db.QueryRow(fmt.Sprintf(`SELECT active FROM event WHERE id = %d`, id))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return false
	}

	var status bool
	query.Scan(&status)

	return status

}
