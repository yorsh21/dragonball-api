package main

import (
	"dragonball-api/pkg/initialize"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start Application")

	db, err := initialize.InitDatabase("./db/sqlite.db")
	if err != nil {
		log.Fatal("Error init database: ", err)
	}
	fmt.Println("Init database success")

	resources := initialize.InitResources(db)
	fmt.Println("Init resources success")

	if err := initialize.InitRouter(resources); err != nil {
		log.Fatal("Error init router: ", err)
	}
	fmt.Println("Init router success")
}
