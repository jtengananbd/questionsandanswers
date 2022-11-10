package entity

import (
	"time"
)

type Question struct {
	ID        string    `json:"id"`
	Tittle    string    `json:"tittle"`
	Statement string    `json:"statement"`
	CreatedOn time.Time `json:"createdOn"`
	Tags      string    `json:"tags"`
}
