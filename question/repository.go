package question

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/errors"
)

type Repository interface {
	Create(question entity.Question) (entity.Question, error)
	GetByID(ID string) (entity.Question, error)
	List(user string) ([]entity.Question, error)
	Update(question entity.Question) (entity.Question, error)
	Delete(ID string) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db}
}

func (r repository) Create(question entity.Question) (entity.Question, error) {
	question.CreatedOn = time.Now()
	err := r.DB.QueryRow(
		"INSERT INTO questions(user_id, tittle, statement, tags, created_on) VALUES($1, $2, $3, $4, $5) RETURNING id, user_id, tittle, statement, tags, created_on",
		question.UserID, question.Tittle, question.Statement, question.Tags, question.CreatedOn).
		Scan(&question.ID, &question.UserID, &question.Tittle, &question.Statement, &question.Tags, &question.CreatedOn)

	if err != nil {
		log.Println("insert question failed", err)
		return entity.Question{}, err
	}

	return question, err
}

func (r repository) Update(question entity.Question) (entity.Question, error) {
	_, err := r.DB.Exec("UPDATE questions SET tittle=$1, statement=$2, tags=$3 WHERE id=$4",
		question.Tittle, question.Statement, question.Tags, question.ID)

	if err != nil {
		log.Println("update question failed", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.Question{}, &errors.ResourceNotFoundError{Resource: "Question", ID: question.ID}
		}
		return entity.Question{}, err
	}
	return question, err
}

func (r repository) GetByID(ID string) (entity.Question, error) {
	question := entity.Question{}
	err := r.DB.QueryRow("SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE id=$1",
		ID).Scan(&question.ID, &question.UserID, &question.Tittle, &question.Statement, &question.Tags, &question.CreatedOn)

	if err != nil {
		log.Println("get question by id failed", err)
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.Question{}, &errors.ResourceNotFoundError{Resource: "Question", ID: ID}
		}
		return entity.Question{}, err
	}

	return question, err
}

func (r repository) Delete(ID string) error {
	_, err := r.DB.Exec("DELETE FROM questions WHERE id=$1", ID)
	return err
}

func (r repository) List(user string) ([]entity.Question, error) {

	var rows *sql.Rows
	var err error

	if user != "" {
		rows, err = r.DB.Query(
			"SELECT id, user_id, tittle, statement, tags, created_on FROM questions WHERE user_id=$1", user)
	} else {
		rows, err = r.DB.Query(
			"SELECT id, user_id, tittle, statement, tags, created_on FROM questions")
	}

	if err != nil {
		log.Println("list question failed", err)
		return []entity.Question{}, err
	}

	defer rows.Close()

	questions := []entity.Question{}

	for rows.Next() {
		var q entity.Question
		if err := rows.Scan(&q.ID, &q.UserID, &q.Tittle, &q.Statement, &q.Tags, &q.CreatedOn); err != nil {
			return []entity.Question{}, err
		}
		questions = append(questions, q)
	}

	return questions, nil
}
