package question

import "github.com/jtengananbd/questionsandanswers/entity"

type Service interface {
	Create(question entity.Question) (entity.Question, error)
	GetByID(ID string) (entity.Question, error)
	List() ([]entity.Question, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Create(question entity.Question) (entity.Question, error) {
	return s.repo.Create(question)
}

func (s service) GetByID(ID string) (entity.Question, error) {
	return s.repo.GetByID(ID)
}

func (s service) List() ([]entity.Question, error) {
	return s.repo.List()
}
