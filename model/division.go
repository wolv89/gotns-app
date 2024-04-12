package model


import (
	"errors"
	"fmt"
	"os"

	"github.com/wolv89/gotnsapp/util"
)


type Division struct {
	Id int			`json:"id"`
	Event int		`json:"event"`
	Name string		`json:"name"`
	Path string		`json:"path"`
	Active bool		`json:"active"`
	Seq int			`json:"seq"`
	Updated string	`json:"updated"`
	Style int		`json:"style"`
	Teams bool		`json:"teams"`
}


type DivisionStatus struct {
	Entrants bool 	`json:"entrants"`
	Matches bool 	`json:"matches"`
}


func divisionInit() {

	dbCreateTable("division", string(`
		CREATE TABLE division (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event INTEGER,
			name TEXT,
			path TEXT,
			active BOOL,
			seq INTEGER,
			updated TEXT,
			style INTEGER,
			teams BOOL
		)
	`))

}



func DivisionCreate(event int, name string, state bool, style int, teams bool) (bool, string) {

	if name == "" {
		return false, "Bad request"
	}

	path := util.StringToPath(name)
	query := db.QueryRow(fmt.Sprintf(`
		SELECT id FROM division
		WHERE event = %d AND
			(name = "%s" OR path = "%s")`, event, name, path))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return false, query.Err().Error()
	}

	var id int
	query.Scan(&id)

	if id > 0 {
		return false, "A division with that name or path already exists"
	}

	query = db.QueryRow(fmt.Sprintf(`
		SELECT COUNT(id) FROM division
		WHERE event = %d`, event))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return false, query.Err().Error()
	}

	var seq int
	query.Scan(&seq)

	seq++

	_, err := db.Query(`
		INSERT INTO division
			(event,name,path,active,seq,updated,style,teams)
		VALUES 
			(?,?,?,?,?,?,?,?)`,
		event, name, path, state, seq, dbNow(), style, teams)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		return false, err.Error()
	}

	return true, path

}


func GetActiveDivisions(event int) ([]Division, error) {

	query, err := db.Query(fmt.Sprintf(`
		SELECT *
		FROM division
		WHERE active = TRUE
		AND event = %d
		ORDER BY seq ASC
	`, event))

	defer query.Close()

	if err != nil {
		return nil, err
	}

	var list []Division

	for query.Next() {

		var div Division

		if qerr := query.Scan(&div.Id, &div.Event, &div.Name, &div.Path, &div.Active, &div.Seq, &div.Updated, &div.Style, &div.Teams); qerr != nil {
			return nil, qerr
		}

		list = append(list, div)

	}

	if len(list) <= 0 {
		return nil, errors.New("none")
	}

	if err = query.Err(); err != nil {
		return nil, err
	}

	return list, nil

}


func GetAllDivisions(event int) ([]Division, error) {

	query, err := db.Query(fmt.Sprintf(`
		SELECT *
		FROM division
		WHERE event = %d
		ORDER BY seq ASC
	`, event))

	defer query.Close()

	if err != nil {
		return nil, err
	}

	var list []Division

	for query.Next() {

		var div Division

		if qerr := query.Scan(&div.Id, &div.Event, &div.Name, &div.Path, &div.Active, &div.Seq, &div.Updated, &div.Style, &div.Teams); qerr != nil {
			return nil, qerr
		}

		list = append(list, div)

	}

	if len(list) <= 0 {
		return nil, errors.New("none")
	}

	if err = query.Err(); err != nil {
		return nil, err
	}

	return list, nil

}


func GetDivisionByPath(event int, path string) (Division, error) {

	var div Division

	if len(path) <= 0 {
		return div, errors.New("No path provided")
	}

	query := db.QueryRow(fmt.Sprintf(`SELECT * FROM division WHERE event = %d AND path = "%s"`, event, path))

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return div, query.Err()
	}

	if qerr := query.Scan(&div.Id, &div.Event, &div.Name, &div.Path, &div.Active, &div.Seq, &div.Updated, &div.Style, &div.Teams); qerr != nil {
		return div, qerr
	}

	return div, nil

}
