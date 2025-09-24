package models

import "time"

type Event struct {
	UserId  int       `json:"user_id"`
	EventId int       `json:"event_id"`
	Date    time.Time `json:"date"`
	Name    string    `json:"name"`
}
