package model


type Division struct {
	id int
	event int
	name string
	active bool
	seq int
	updated string
	class int
}


func divisionInit() {

	dbCreateTable("division", string(`
		CREATE TABLE division (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event INTEGER,
			name TEXT,
			active BOOL,
			seq INTEGER,
			updated TEXT,
			class INTEGER
		)
	`))

}