package question

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jtengananbd/questionsandanswers/entity"
	customerr "github.com/jtengananbd/questionsandanswers/errors"
	"github.com/stretchr/testify/assert"
)

var usrId = "user@mail.com"
var questionMock = &entity.Question{
	UserID:    &usrId,
	Tittle:    "How to fix error ...",
	Statement: "I'm having an issue when ...",
	CreatedOn: time.Now(),
	Tags:      "golang, go, panic, error",
}
var answerMock = &entity.Answer{
	ID:         "1",
	QuestionID: 1,
	UserID:     "user@mail.com",
	Comment:    "coment",
	CreatedOn:  time.Now(),
}

func TestQuestionService_Create(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	res, err := service.Create(*questionMock)

	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, "1", res.ID)
}

func TestQuestionService_CreateWithError(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	fmt.Println(questionMock)
	q1 := *questionMock
	q1.Statement = "error"

	_, err := service.Create(q1)
	assert.Error(t, err)
	assert.Equal(t, "some database error", err.Error())
}

func TestQuestionService_Update(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	q1 := *questionMock
	q1.Statement = "An updated statement"
	res, err := service.Update(q1)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "An updated statement", res.Statement)
}

func TestQuestionService_UpdateNotFound(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	q1 := *questionMock
	q1.ID = "0"
	res, err := service.Update(q1)

	assert.Error(t, err)
	assert.Equal(t, "Resource Question with ID:0 not found", err.Error())
	assert.Equal(t, entity.Question{}, res)
}

func TestQuestionService_UpdateFails(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	q1 := *questionMock
	q1.Statement = "error"
	res, err := service.Update(q1)

	assert.Error(t, err)
	assert.Equal(t, "some database error", err.Error())
	assert.Equal(t, entity.Question{}, res)
}

func TestQuestionService_UpdateWithAnswer(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	q1 := *questionMock
	q1.ID = "1"
	q1.Statement = "An updated statement"

	a1 := *answerMock
	q1.Answer = &a1

	res, err := service.Update(q1)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "An updated statement", res.Statement)
}

func TestQuestionService_UpdateWithAnswerFails(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	q1 := *questionMock
	q1.ID = "100"
	q1.Statement = "An updated statement"

	a1 := *answerMock
	q1.Answer = &a1

	res, err := service.Update(q1)
	//TODO
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "An updated statement", res.Statement)
}

type mockRepo struct{}

func (m mockRepo) Create(question entity.Question) (entity.Question, error) {
	if question.Statement == "error" {
		return entity.Question{}, errors.New("some database error")
	}
	question.ID = "1"
	return question, nil
}

func (m mockRepo) GetByID(ID string) (entity.Question, error) {
	switch ID {
	case "0":
		return entity.Question{}, &customerr.ResourceNotFoundError{Resource: "Question", ID: ID}
	default:
		q.ID = ID
		return *q, nil
	}
}

func (m mockRepo) List(user string) ([]entity.Question, error) {
	return []entity.Question{}, nil
}

func (m mockRepo) Update(question entity.Question) (entity.Question, error) {
	if question.Statement == "error" {
		return entity.Question{}, errors.New("some database error")
	}
	return question, nil
}

func (m mockRepo) Delete(ID string) error {
	return nil
}

type answerRepoMock struct {
}

func (am answerRepoMock) Create(answer entity.Answer) (entity.Answer, error) {
	return answer, nil
}

func (am answerRepoMock) Update(answer entity.Answer) (entity.Answer, error) {
	if answer.Comment == "error" {
		return entity.Answer{}, errors.New("some database error")
	}
	return entity.Answer{}, nil
}

func (am answerRepoMock) GetByQuestionID(ID string) (entity.Answer, error) {
	if ID == "1" {
		return *answerMock, nil
	}
	if ID == "100" {
		return entity.Answer{}, &customerr.ResourceNotFoundError{Resource: "Answer", ID: ID}
	}
	return entity.Answer{}, nil
}

func (am answerRepoMock) DeleteByQuestionID(ID string) error {
	return nil
}
