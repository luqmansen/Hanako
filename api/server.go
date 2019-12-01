package api

import (
	"github.com/joho/godotenv"
	"github.com/luqmansen/hanako/api/controllers"
	"os"
)

var server = controllers.Server{}

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func Run() {

	dbUsername := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbDriver := os.Getenv("db_driver")
	dbPort := os.Getenv("db_port")

	server.Initialize(dbDriver, dbUsername, dbPassword, dbPort, dbHost, dbName)
	server.Run(":" + os.Getenv("PORT"))

}
