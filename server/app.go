package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jtengananbd/questionsandanswers/answer"
	"github.com/jtengananbd/questionsandanswers/constant"
	"github.com/jtengananbd/questionsandanswers/question"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func (server *Server) Init(user string, password string, dbname string) {

	connectionValues := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	server.DB, err = sql.Open("postgres", connectionValues)
	if err != nil {
		log.Fatal(err)
	}
	checkError(err)

	// check db
	err = server.DB.Ping()
	checkError(err)
	fmt.Println("Connected to DB!")
	server.initDBTables()

	server.Router = mux.NewRouter().StrictSlash(true)

	server.initializeRoutes()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initDBTables() {
	if _, err := server.DB.Exec(constant.CreateQuestionsQuery); err != nil {
		log.Fatal("initTable Questions ", err)
	}
	if _, err := server.DB.Exec(constant.CreateAnswersQuery); err != nil {
		log.Fatal("initTable Answers ", err)
	}
}

func (server *Server) StartServer(address string) {
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", question.Home)

	questionAPI := question.NewAPI(question.NewService(question.NewRepository(server.DB), answer.NewRepository(server.DB)))

	server.Router.HandleFunc("/questions", questionAPI.Create).Methods("POST")
	server.Router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.GetByID).Methods("GET")
	server.Router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.Update).Methods("PUT")
	server.Router.HandleFunc("/questions", questionAPI.List).Methods("GET")
	server.Router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.Delete).Methods("DELETE")
}
