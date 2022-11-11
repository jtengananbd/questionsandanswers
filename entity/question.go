package entity

import (
	"time"
)

type Question struct {
	ID        string    `json:"id"`
	UserID    *string   `json:"userID"`
	Tittle    string    `json:"tittle"`
	Statement string    `json:"statement"`
	CreatedOn time.Time `json:"createdOn"`
	Tags      string    `json:"tags"`
	Answer    *Answer   `json:"answer,omitempty"`
}
