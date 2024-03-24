package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/joho/godotenv"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/router"

)



func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load local env %s", err)
		os.Exit(1)
	}

	fmt.Println("## Server Init")

	model.DatabaseInit()
	defer model.DatabaseClose()

	fmt.Println("## Database connected")

	router := router.LoadRouter()

	fmt.Println("## Routes loaded")

	server := http.Server {
		Addr:		":8040",
		Handler:	router,
	}

	fmt.Println("## Server Started")

	server.ListenAndServe()

}
