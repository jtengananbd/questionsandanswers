package question

import (
	"log"
	"strconv"

	"github.com/jtengananbd/questionsandanswers/answer"
	"github.com/jtengananbd/questionsandanswers/entity"
	"github.com/jtengananbd/questionsandanswers/errors"
)

type Service interface {
	Create(question entity.Question) (entity.Question, error)
	Update(question entity.Question) (entity.Question, error)
	GetByID(ID string) (entity.Question, error)
	List(user string) ([]entity.Question, error)
	Delete(ID string) error
}

type service struct {
	repo             Repository
	answerRepository answer.Repository
}

func NewService(repo Repository, answerRepository answer.Repository) Service {
	return service{repo: repo, answerRepository: answerRepository}
}

func (s service) Create(question entity.Question) (entity.Question, error) {
	return s.repo.Create(question)
}

func (s service) Update(question entity.Question) (entity.Question, error) {
	questionDB, err := s.repo.GetByID(question.ID)

	if err != nil {
		log.Println("update question failed to get existing question", err)
		return entity.Question{}, err
	}

	questionDB.Statement = question.Statement
	questionDB.Tittle = question.Tittle
	questionDB.Tags = question.Tags

	q, err := s.repo.Update(questionDB)
	if err != nil {
		log.Println("update question failed", err)
		return entity.Question{}, err
	}

	if question.Answer != nil {

		answerDB, err := s.answerRepository.GetByQuestionID(questionDB.ID)
		if err != nil {

			switch err.(type) {
			case *errors.ResourceNotFoundError:
				qid, _ := strconv.Atoi(questionDB.ID)
				(*question.Answer).QuestionID = qid
				answerDB, err := s.answerRepository.Create(*question.Answer)
				if err != nil {
					log.Println("update question failed to create new answer", err)
					return entity.Question{}, err
				}
				q.Answer = &answerDB
				return q, nil
			default:
				return entity.Question{}, err
			}
		}
		answerDB.Comment = question.Answer.Comment

		answerDB, err = s.answerRepository.Update(answerDB)
		if err != nil {
			log.Println("update question failed to update answer", err)
			return entity.Question{}, err
		}
		q.Answer = &answerDB
	}
	return q, nil
}

func (s service) GetByID(ID string) (entity.Question, error) {
	questionDB, err := s.repo.GetByID(ID)
	if err != nil {
		return entity.Question{}, err
	}
	answerDB, err := s.answerRepository.GetByQuestionID(questionDB.ID)
	if err != nil {
		switch err.(type) {
		case *errors.ResourceNotFoundError:
			questionDB.Answer = nil
		}
	} else {
		questionDB.Answer = &answerDB
	}
	return questionDB, nil
}

func (s service) Delete(ID string) error {
	_, err := s.answerRepository.GetByQuestionID(ID)
	if err != nil {
		return err
	}
	s.answerRepository.DeleteByQuestionID(ID)
	return s.repo.Delete(ID)
}

func (s service) List(user string) ([]entity.Question, error) {
	return s.repo.List(user)
}
