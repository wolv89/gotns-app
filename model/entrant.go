package model


type Entrant struct {
	id int
	player1 int
	player2 int
	isteam bool
	seed int
}


func entrantInit() {

	dbCreateTable("entrant", string(`
		CREATE TABLE entrant (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			player1 INTEGER,
			player2 INTEGER,
			isteam BOOL,
			seed INTEGER
		)
	`))

}