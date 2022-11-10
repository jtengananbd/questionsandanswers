package main

//"github.com/jtengananbd/questionsandanswers/pkg"
import (
	"github.com/jtengananbd/questionsandanswers/server"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mypassword"
	dbname   = "postgres"
)

func main() {

	server := server.Server{}

	server.Init(user, password, dbname)

	defer server.DB.Close()
	server.StartServer(":8080")

}