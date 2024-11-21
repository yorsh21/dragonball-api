package main

import (
	"dragonball-api/db"
	"dragonball-api/pkg/initialize"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start Migrations")

	conn, err := initialize.InitDatabase("./db/sqlite.db")
	if err != nil {
		log.Fatal("Error init database: ", err)
	}
	fmt.Println("Init database success")

	migration := db.NewMigration(conn)
	migration.Character()

}
