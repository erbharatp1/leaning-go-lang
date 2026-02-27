package model

import (
	"damo-go/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `json:"name" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	Description string    `json:"description" binding:"required"`
	UserId      string
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, location, description, UserId, DateTime) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()

	if e.DateTime.IsZero() {
		e.DateTime = time.Now()
	}

	result, err := stmt.Exec(e.Name, e.Location, e.Description, e.UserId, e.DateTime)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return err
	}

	e.ID = id
	return nil
}

func EventsList() []Event {
	query := "SELECT id, name, location, description, UserId, DateTime FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying events: %v", err)
		return []Event{}
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Location, &e.Description, &e.UserId, &e.DateTime)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		events = append(events, e)
	}
	return events
}
