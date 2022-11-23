package question

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/test"
)

var userID = "user@mail.com"
var q = &entity.Question{
	ID:        "1",
	UserID:    &userID,
	Tittle:    "How to fix error ...",
	Statement: "I'm having an issue when ...",
	CreatedOn: time.Now(),
	Tags:      "golang, go, panic, error",
}

func TestQuestionRepository_GetByID(t *testing.T) {
	//https://medium.easyread.co/unit-test-sql-in-golang-5af19075e68e
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE id=$1"

	rows := sqlmock.NewRows([]string{"id", "user_id", "tittle", "statement", "tags", "created_on"}).
		AddRow(q.ID, q.UserID, q.Tittle, q.Statement, q.Tags, q.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.ID).WillReturnRows(rows)

	answer, err := repo.GetByID(q.ID)

	assert.NoError(t, err)
	assert.NotNil(t, answer)
}

func TestQuestionRepository_GetByIDFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE id=$1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.ID).WillReturnError(errors.New("Unexpected Error"))

	question, err := repo.GetByID(q.ID)

	assert.Error(t, err)
	assert.Empty(t, question.ID)
}

func TestQuestionRepository_Create(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "INSERT INTO questions(user_id, tittle, statement, tags, created_on) VALUES($1, $2, $3, $4, $5) RETURNING id, user_id, tittle, statement, tags, created_on"

	rows := sqlmock.NewRows([]string{"id", "user_id", "tittle", "statement", "tags", "created_on"}).
		AddRow(q.ID, q.UserID, q.Tittle, q.Statement, q.Tags, q.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.UserID, q.Tittle, q.Statement, q.Tags, test.AnyTime{}).
		WillReturnRows(rows)

	question, err := repo.Create(*q)

	assert.NoError(t, err)
	assert.Equal(t, q.ID, question.ID)
}

func TestQuestionRepository_CreateFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "INSERT INTO questions(user_id, tittle, statement, tags, created_on) VALUES($1, $2, $3, $4, $5) RETURNING id, user_id, tittle, statement, tags, created_on"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.UserID, q.Tittle, q.Statement, q.Tags, test.AnyTime{}).
		WillReturnError(errors.New("Unspected Error"))

	question, err := repo.Create(*q)

	assert.Error(t, err)
	assert.Empty(t, question.ID)
}

func TestQuestionRepository_Update(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "UPDATE questions SET tittle=$1, statement=$2, tags=$3 WHERE id=$4"
	qt := *q
	qt.Tittle = "updated tittle"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(qt.Tittle, qt.Statement, qt.Tags, qt.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	question, err := repo.Update(qt)

	assert.NoError(t, err)
	assert.Equal(t, qt.Tittle, question.Tittle)
}

func TestQuestionRepository_UpdateFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "UPDATE questions SET tittle=$1, statement=$2, tags=$3 WHERE id=$4"
	qt := *q
	qt.Tittle = "updated tittle"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(qt.Tittle, qt.Statement, qt.Tags, qt.ID).
		WillReturnError(errors.New("Unspected Error"))

	question, err := repo.Update(qt)

	assert.Error(t, err)
	assert.Empty(t, question.ID)

}

func TestQuestionRepository_UpdateNotFound(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "UPDATE questions SET tittle=$1, statement=$2, tags=$3 WHERE id=$4"
	qt := *q
	qt.Tittle = "updated tittle"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(qt.Tittle, qt.Statement, qt.Tags, qt.ID).
		WillReturnError(errors.New("no rows in result set"))

	question, err := repo.Update(qt)

	assert.Error(t, err)
	assert.Empty(t, question.ID)

}

func TestQuestionRepository_Delete(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "DELETE FROM questions WHERE id=$1"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(q.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(q.ID)

	assert.NoError(t, err)
}

func TestRepository_DeleteFails(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "DELETE FROM questions WHERE id=$1"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(q.ID).WillReturnError(errors.New("Unspected Error"))

	err := repo.Delete(q.ID)

	assert.Error(t, err)
}

func TestQuestionRepository_List(t *testing.T) {
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE user_id=$1"

	rows := sqlmock.NewRows([]string{"id", "user_id", "tittle", "statement", "tags", "created_on"}).
		AddRow(q.ID, q.UserID, q.Tittle, q.Statement, q.Tags, q.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.UserID).WillReturnRows(rows)

	answers, err := repo.List(*q.UserID)

	assert.NoError(t, err)
	assert.NotEmpty(t, answers)

}
