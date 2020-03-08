package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/luqmansen/hanako/api/models/postgres"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database \n ", Dbdriver)
		panic(err)
	} else {
		fmt.Printf("We are connected to the %s database \n", Dbdriver)

		server.DB.Debug().AutoMigrate(&postgres.Account{}, &postgres.Anime2{}) //database migration
		server.Router = mux.NewRouter()
		server.initializeRoutes()
	}
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
