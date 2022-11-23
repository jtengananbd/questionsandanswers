package answer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/test"
)

var a = &entity.Answer{
	ID:         "1",
	QuestionID: 2,
	UserID:     "user@mail.com",
	Comment:    "coment",
	CreatedOn:  time.Now(),
}

func TestRepository_GetByQuestionID(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, question_id, user_id, comment, created_on FROM answers WHERE question_id=$1"

	rows := sqlmock.NewRows([]string{"id", "question_id", "user_id", "comment", "created_on"}).
		AddRow(a.ID, a.QuestionID, a.UserID, a.Comment, a.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(a.ID).WillReturnRows(rows)

	answer, err := repo.GetByQuestionID(a.ID)

	assert.NotNil(t, answer)
	assert.NoError(t, err)
}

func TestRepository_GetByQuestionIDFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, question_id, user_id, comment, created_on FROM answers WHERE question_id=$1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(a.ID).WillReturnError(errors.New("Unexpected Error"))

	answer, err := repo.GetByQuestionID(a.ID)

	assert.Error(t, err)
	assert.Empty(t, answer.ID)
}

func TestRepository_Create(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "INSERT INTO answers(question_id, user_id, comment, created_on) VALUES($1, $2, $3, $4) RETURNING id, question_id, user_id, comment, created_on"

	rows := sqlmock.NewRows([]string{"id", "question_id", "user_id", "comment", "created_on"}).
		AddRow(a.ID, a.QuestionID, a.UserID, a.Comment, a.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(a.QuestionID, a.UserID, a.Comment, test.AnyTime{}).WillReturnRows(rows)

	answer, err := repo.Create(*a)

	assert.NoError(t, err)
	assert.Equal(t, a.ID, answer.ID)
}

func TestRepository_CreateFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "INSERT INTO answers(question_id, user_id, comment, created_on) VALUES($1, $2, $3, $4) RETURNING id, question_id, user_id, comment, created_on"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(a.QuestionID, a.UserID, a.Comment, a.CreatedOn).WillReturnError(errors.New("Unexpected Error"))

	answer, err := repo.Create(*a)

	assert.Error(t, err)
	assert.Empty(t, answer.ID)
}

func TestRepository_Update(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "UPDATE answers SET comment=$1 WHERE id=$2"
	ans := *a
	ans.Comment = "updated comment"
	fmt.Println(a)
	fmt.Println(ans)
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(ans.Comment, a.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	answer, err := repo.Update(ans)

	assert.NoError(t, err)
	assert.Equal(t, ans.Comment, answer.Comment)
}

func TestRepository_UpdateFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "UPDATE answers SET comment=$1 WHERE id=$2"
	ans := a
	ans.Comment = "updated comment"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(a.Comment, a.ID).WillReturnError(errors.New("Unspected Error"))

	answer, err := repo.Update(*a)

	assert.Error(t, err)
	assert.Empty(t, answer.ID)
}

func TestRepository_Delete(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "DELETE FROM answers WHERE question_id=$1"

	ID := strconv.Itoa(a.QuestionID)

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteByQuestionID(ID)

	assert.NoError(t, err)
}

func TestRepository_DeleteFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "DELETE FROM answers WHERE question_id=$1"

	ID := strconv.Itoa(a.QuestionID)

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(ID).WillReturnError(errors.New("Unspected Error"))

	err := repo.DeleteByQuestionID(ID)

	assert.Error(t, err)
}
