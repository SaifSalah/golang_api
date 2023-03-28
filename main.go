package main

import (
	"log"
)

func main() {

	dbConn, err := PostgresConn()

	if err != nil {
		log.Fatal(err)
	}

	if err := dbConn.Init(); err != nil {
		log.Fatal(err)
	}

	server := newAPIServer(":1992", dbConn)
	server.Run()
}
