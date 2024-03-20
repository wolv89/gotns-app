package model


type Player struct {
	id int
	firstname string
	lastname string
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