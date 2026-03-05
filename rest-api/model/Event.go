package model

import (
	"leaning-go-lang/db"
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

func (e *Event) Update() error {
	query := `UPDATE events SET name = ?, location = ?, description = ?, UserId = ?, DateTime = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing update statement: %v", err)
		return err
	}
	defer stmt.Close()

	if e.DateTime.IsZero() {
		e.DateTime = time.Now()
	}

	_, err = stmt.Exec(e.Name, e.Location, e.Description, e.UserId, e.DateTime, e.ID)
	if err != nil {
		log.Printf("Error executing update statement: %v", err)
		return err
	}
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

func FindByName(name string) (*Event, error) {
	log.Println("FindByName")
	query := "SELECT id, name, location, description, UserId, DateTime FROM events WHERE name = ?"
	row := db.DB.QueryRow(query, name)
	log.Println(query)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.UserId, &event.DateTime)
	if err != nil {
		log.Printf("Error scanning event by name: %v", err)
		return nil, err
	}
	log.Println(event)
	return &event, nil
}

func FindByID(id int64) (*Event, error) {
	log.Printf("FindByID: %d", id)
	query := "SELECT id, name, location, description, UserId, DateTime FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.UserId, &event.DateTime)
	if err != nil {
		log.Printf("Error scanning event by ID: %v", err)
		return nil, err
	}
	return &event, nil
}
