package model


type Division struct {
	Id int			`json:"id"`
	Event int		`json:"event"`
	Name string		`json:"name"`
	Path string		`json:"path"`
	Active bool		`json:"active"`
	Seq int			`json:"seq"`
	Updated string	`json:"updated"`
	Class int		`json:"class"`
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
			class INTEGER
		)
	`))

}