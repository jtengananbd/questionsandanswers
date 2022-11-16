package question

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/test"
	"github.com/stretchr/testify/assert"
)

var usr = "julio@mail.com"
var q3 = entity.Question{
	UserID: &usr, Tittle: "tittle1", Statement: "how to...?", Tags: "go, golang, REST",
}

func TestAPI(t *testing.T) {

	router := mux.NewRouter()

	questionAPI := NewAPI(questionServiceMock{})
	router.HandleFunc("/questions", questionAPI.Create).Methods("POST")
	router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.GetByID).Methods("GET")
	router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.Update).Methods("PUT")
	router.HandleFunc("/questions", questionAPI.List).Methods("GET")
	router.HandleFunc("/questions/{id:[0-9]+}", questionAPI.Delete).Methods("DELETE")

	payload := []byte(`{"userID": "julio@mail.com", "tittle": "tittle1", "statement": "how to...?","tags": "go, golang, REST"}`)
	req, _ := http.NewRequest("POST", "/questions", bytes.NewBuffer(payload))
	response := test.ExecuteRequest(req, router)
	assert.NotNil(t, response.Body)
	var questionRes entity.Question
	json.Unmarshal(response.Body.Bytes(), &questionRes)
	assert.NotNil(t, questionRes)
	test.CheckResponseCode(t, http.StatusCreated, response.Code)

	tests := []test.APITestCase{
		{Name: "create Question", Method: "POST", URL: "/questions", Body: `{"userID": "julio@mail.com", "tittle": "tittle1", "statement": "how to...?","tags": "go, golang, REST"}`, WantStatus: http.StatusCreated, WantResponse: `*tittle1*`},
		{Name: "Get Question", Method: "GET", URL: "/questions/1", Body: "", WantStatus: http.StatusOK, WantResponse: `*tittle1*`},
		{Name: "Update Question", Method: "PUT", URL: "/questions/2", Body: `{"userID": "julio@mail.com", "tittle": "tittle updated", "statement": "how to...?","tags": "go, golang, REST"}`, WantStatus: http.StatusOK, WantResponse: `*tittle updated*`},
		{Name: "Delete Question", Method: "DELETE", URL: "/questions/2", Body: "", WantStatus: http.StatusOK},
		{Name: "List Question", Method: "GET", URL: "/questions", Body: "", WantStatus: http.StatusOK, WantResponse: `*tittle1*`},
	}

	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}

}

type questionServiceMock struct{}

func (s questionServiceMock) Create(question entity.Question) (entity.Question, error) {
	return q3, nil
}
func (s questionServiceMock) Update(question entity.Question) (entity.Question, error) {
	tmp := q3
	tmp.Tittle = "tittle updated"
	return tmp, nil
}
func (s questionServiceMock) GetByID(ID string) (entity.Question, error) {
	return q3, nil
}
func (s questionServiceMock) List(user string) ([]entity.Question, error) {
	return []entity.Question{q3}, nil
}
func (s questionServiceMock) Delete(ID string) error {
	return nil
}
