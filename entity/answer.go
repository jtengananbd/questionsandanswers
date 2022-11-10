package entity

import "time"

type Answer struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"questionId"`
	Comment    string    `json:"comment"`
	CreatedOn  time.Time `json:"createdOn"`
}
