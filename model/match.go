package model


type Match struct {
	id int
	division int
	entrant1 int
	entrant2 int
	score string
	notes string
	seq int
	start string
	updated string
	status int
	winner int
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