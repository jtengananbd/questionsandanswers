package entity

import "time"

type Answer struct {
	ID         string    `json:"id"`
	QuestionID int       `json:"questionID"`
	UserID     string    `json:"userID"`
	Comment    string    `json:"comment"`
	CreatedOn  time.Time `json:"createdOn"`
}
