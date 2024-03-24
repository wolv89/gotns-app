package model


type Entrant struct {
	Id int			`json:"id"`
	Player1 int		`json:"player1"`
	Player2 int		`json:"player2"`
	IsTeam bool		`json:"isteam"`
	Seed int		`json:"seed"`
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