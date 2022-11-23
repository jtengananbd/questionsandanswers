package question

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/stretchr/testify/assert"
)

var usrId = "user@mail.com"
var question = &entity.Question{
	UserID:    &usrId,
	Tittle:    "How to fix error ...",
	Statement: "I'm having an issue when ...",
	CreatedOn: time.Now(),
	Tags:      "golang, go, panic, error",
}

func TestService_CreateQuestion(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	questionDB := *question
	questionDB.ID = "1"
	res, err := service.Create(questionDB)

	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, "1", res.ID)
}

func TestService_CreateQuestionWithError(t *testing.T) {

	service := NewService(mockRepo{}, answerRepoMock{})
	fmt.Println(question)
	q1 := *question
	q1.Statement = "error"

	_, err := service.Create(q1)
	assert.Error(t, err)
	assert.Equal(t, "some database error", err.Error())
}

type mockRepo struct{}

func (m mockRepo) Create(question entity.Question) (entity.Question, error) {
	if question.Statement == "error" {
		return entity.Question{}, errors.New("some database error")
	}
	return question, nil
}

func (m mockRepo) GetByID(ID string) (entity.Question, error) {
	return entity.Question{}, nil
}

func (m mockRepo) List(user string) ([]entity.Question, error) {
	return []entity.Question{}, nil
}

func (m mockRepo) Update(question entity.Question) (entity.Question, error) {
	return entity.Question{}, nil
}

func (m mockRepo) Delete(ID string) error {
	return nil
}

type answerRepoMock struct {
}

func (am answerRepoMock) Create(answer entity.Answer) (entity.Answer, error) {
	return entity.Answer{}, nil
}

func (am answerRepoMock) Update(answer entity.Answer) (entity.Answer, error) {
	return entity.Answer{}, nil
}

func (am answerRepoMock) GetByQuestionID(ID string) (entity.Answer, error) {
	return entity.Answer{}, nil
}

func (am answerRepoMock) DeleteByQuestionID(ID string) error {
	return nil
}
