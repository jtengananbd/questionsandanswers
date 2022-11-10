package question

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jtengananbd/questionsandanswers/entity"
)

type API interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type api struct {
	serv Service
}

func NewAPI(service Service) API {
	return api{service}
}

func (a api) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please sent a valid question")
	}

	var questionRq entity.Question
	json.Unmarshal(body, &questionRq)

	createdEntity, err := a.serv.Create(questionRq)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEntity)
}

func (api api) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	question, err := api.serv.GetByID(id)

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (api api) List(w http.ResponseWriter, r *http.Request) {
	questions, err := api.serv.List()

	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
