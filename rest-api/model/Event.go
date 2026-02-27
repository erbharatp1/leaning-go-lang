package model

import "time"

type Event struct {
	ID          string
	Name        string `json:"name" binding:"required"`
	Location    string `json:"location" binding:"required"`
	DateTime    time.Time
	Description string `json:"description" binding:"required"`
	UserId      string
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func EventsList() []Event {
	return events
}
