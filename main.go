package main

import (
	"log"

	"github.com/steallers/server-utility/database"
)

type Fm struct {
	ID   string
	Name string
}

func (Fm) TableName() string {
	return "maindb.fm"
}
func main() {
	// Running Services
	dbConn := database.StartDatabaseService()
	log.Println(dbConn)
	which := database.DatabaseService{}
	ndr := Fm{}
	log.Println(which.GetConnectors().AutoMigrate(&ndr))
	log.Println("connector status: ", which.GetConnectors())

	// Running Repository
}
