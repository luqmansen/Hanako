package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/luqmansen/hanako/api/controllers"
	"log"
	"os"
)

var server = controllers.Server{}

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print("Error : ", e)
	}
}
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		dbUsername := os.Getenv("db_user")
		dbPassword := os.Getenv("db_pass")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")
		dbDriver := os.Getenv("db_driver")
		dbPort := os.Getenv("db_port")

		server.Initialize(dbDriver, dbUsername, dbPassword, dbPort, dbHost, dbName)
		server.Run(":" + os.Getenv("PORT"))
	}
}
