package model

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var db *sql.DB


func DatabaseLaunch() {
	eventInit()
	divisionInit()
	playerInit()
	entrantInit()
	matchInit()
}

func DatabaseNuke() {
	dbDropTable("event")
	dbDropTable("division")
	dbDropTable("player")
	dbDropTable("entrant")
	dbDropTable("match")
}


func DatabaseInit() {

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load local env %s", err)
		os.Exit(1)
	}

	dburl := fmt.Sprintf("libsql://%v.turso.io?authToken=%v", os.Getenv("DB_NAME"), os.Getenv("DB_KEY"))

	db, err = sql.Open("libsql", dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open db %s: %s", dburl, err)
		os.Exit(1)
	}

}


func dbCreateTable(name string, schema string) {

	if dbTableExists(name) {
		fmt.Println("  > [", name, "] table confirmed")
		return
	}

	fmt.Println("  > Creating [", name, "] table")

	_, err := db.Exec(schema)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}

}


func dbDropTable(name string) {

	if !dbTableExists(name) {
		fmt.Println("-- [", name, "] table does not exist")
		return
	}

	fmt.Println("-- Dropping [", name, "] table")

	_, err := db.Exec(fmt.Sprintf("DROP TABLE '%s'", name))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute query: %v\n", err)
		os.Exit(1)
	}

}


func dbTableExists(tablename string) bool {

	query := db.QueryRow(fmt.Sprintf("SELECT name FROM sqlite_master WHERE type = 'table' AND name = '%s'", tablename))
	foundname := ""

	if query.Err() != nil {
		fmt.Println(query.Err().Error())
		return false
	}

	query.Scan(&foundname)
	if foundname != "" {
		return true
	}

	return false

}


func dbNow() string {
	return time.Now().Format(time.DateTime)
}


func DatabaseClose() {
	db.Close()
}
