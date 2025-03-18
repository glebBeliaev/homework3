package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Discription string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}
