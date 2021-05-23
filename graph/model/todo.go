package model

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	UserID    string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}
