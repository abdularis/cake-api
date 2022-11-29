package main

import (
	"cake-api/config"
	"cake-api/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	cfg := config.Get()
	db, err := utils.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) <= 1 {
		log.Fatal("Please provide path to init scheme migration sql file\n" +
			"\tExample:\n" +
			"\t\tgo run migrations/init.go migrations/mysql/cake_service/001-init_scheme.sql")
	}

	initSchemePath := os.Args[1]
	fmt.Printf("Executing %s\n", initSchemePath)

	sqlScript, err := os.ReadFile(initSchemePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(sqlScript))
	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.WithError(err).Fatal("Error while executing init sql scheme")
	}
	fmt.Println("Init scheme success")
}
