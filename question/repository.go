package question

import (
	"database/sql"
	"log"
	"time"

	"github.com/jtengananbd/questionsandanswers/entity"
)

type Repository interface {
	Create(question entity.Question) (entity.Question, error)
	GetByID(ID string) (entity.Question, error)
	List() ([]entity.Question, error)
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
		"INSERT INTO questions(tittle, statement, tags) VALUES($1, $2, $3) RETURNING id, tittle, statement, tags",
		question.Tittle, question.Statement, question.Tags).
		Scan(&question.ID, &question.Tittle, &question.Statement, &question.Tags)

	if err != nil {
		log.Println("insert question failed", err)
		return entity.Question{}, err
	}

	return question, nil
}

func (r repository) GetByID(ID string) (entity.Question, error) {
	question := entity.Question{}
	err := r.DB.QueryRow("SELECT id, tittle, statement, tags FROM questions WHERE id=$1",
		ID).Scan(&question.ID, &question.Tittle, &question.Statement, &question.Tags)

	if err != nil {
		log.Println("get question by id failed", err)
		return entity.Question{}, err
	}

	return question, nil
}

func (r repository) List() ([]entity.Question, error) {
	rows, err := r.DB.Query(
		"SELECT id, tittle, statement, tags FROM questions")

	if err != nil {
		log.Println("list question failed", err)
		return nil, err
	}

	defer rows.Close()

	questions := []entity.Question{}

	for rows.Next() {
		var q entity.Question
		if err := rows.Scan(&q.ID, &q.Tittle, &q.Statement, &q.Tags); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}

	return questions, nil
}
