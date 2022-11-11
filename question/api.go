package question

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/errors"
)

type API interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEntity)
}

func (api api) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please sent a valid question")
	}

	var questionRq entity.Question
	json.Unmarshal(body, &questionRq)
	questionRq.ID = id

	question, err := api.serv.Update(questionRq)

	if err != nil {
		log.Println("question API: error trying to update question ID: ", id, err)
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (api api) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.serv.Delete(id)

	if err != nil {
		log.Println("question API: error trying to delete question ID: ", id, err)
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "deleted success"}`))
}

func (api api) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	question, err := api.serv.GetByID(id)

	if err != nil {
		handleError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

func (api api) List(w http.ResponseWriter, r *http.Request) {

	user := r.URL.Query().Get("user")

	fmt.Println("user:", user)

	questions, err := api.serv.List(user)

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		switch e := err.(type) {
		case *errors.ResourceNotFoundError:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		default:
			log.Println(e)
		}
	}
}
