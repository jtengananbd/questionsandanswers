package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jtengananbd/questionsandanswers/server"
)

/*const (
	host     = "postgresdb"
	port     = 5432
	user     = "postgres"
	password = "mypassword"
	dbname   = "postgres"
)*/

func init() {
	//loads values from .env into the system
	fmt.Println("Loading env vars...")
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}
func main() {

	server := server.Server{}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	server.Init(user, password, dbname, host)

	defer server.DB.Close()
	server.StartServer(":8080")

}
