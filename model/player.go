package model


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