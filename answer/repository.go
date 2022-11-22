package answer

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/errors"
)

type Repository interface {
	Create(answer entity.Answer) (entity.Answer, error)
	Update(answer entity.Answer) (entity.Answer, error)
	GetByQuestionID(ID string) (entity.Answer, error)
	DeleteByQuestionID(ID string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Create(answer entity.Answer) (entity.Answer, error) {
	answer.CreatedOn = time.Now()
	err := r.DB.QueryRow(
		"INSERT INTO answers(question_id, user_id, comment, created_on) VALUES($1, $2, $3, $4) RETURNING id, question_id, user_id, comment, created_on",
		answer.QuestionID, answer.UserID, answer.Comment, answer.CreatedOn).
		Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Comment, &answer.CreatedOn)

	if err != nil {
		log.Println("insert question failed", err)
		return entity.Answer{}, err
	}

	return answer, nil
}

func (r repository) GetByQuestionID(ID string) (entity.Answer, error) {
	answer := entity.Answer{}
	err := r.DB.QueryRow("SELECT id, question_id, user_id, comment, created_on FROM answers WHERE question_id=$1",
		ID).Scan(&answer.ID, &answer.QuestionID, &answer.UserID, &answer.Comment, &answer.CreatedOn)

	if err != nil {
		log.Println("get answer by question_id failed", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.Answer{}, &errors.ResourceNotFoundError{Resource: "Answer", ID: ID}
		}
		return entity.Answer{}, err
	}
	return answer, nil
}

func (r repository) Update(answer entity.Answer) (entity.Answer, error) {
	_, err := r.DB.Exec("UPDATE answers SET comment=$1 WHERE id=$2",
		answer.Comment, answer.ID)

	if err != nil {
		log.Println("update answer failed", err)
		return entity.Answer{}, err
	}
	return answer, nil
}

func (r repository) DeleteByQuestionID(ID string) error {
	_, err := r.DB.Exec("DELETE FROM answers WHERE question_id=$1", ID)
	return err
}
