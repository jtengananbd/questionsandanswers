package question

import (
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

func TestRepository_GetByID(t *testing.T) {
	//https://medium.easyread.co/unit-test-sql-in-golang-5af19075e68e
	db, mock := test.NewMockDB()

	repo := NewRepository(db)

	query := "SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE id=$1"

	rows := sqlmock.NewRows([]string{"id", "user_id", "tittle", "statement", "tags", "created_on"}).
		AddRow(q.ID, q.UserID, q.Tittle, q.Statement, q.Tags, q.CreatedOn)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(q.ID).WillReturnRows(rows)

	answer, err := repo.GetByID(q.ID)

	assert.NotNil(t, answer)
	assert.NoError(t, err)
}
