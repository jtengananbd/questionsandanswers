package answer

import (
	"regexp"
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
