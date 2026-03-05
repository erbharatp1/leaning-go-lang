package model

import (
	"leaning-go-lang/db"
	"leaning-go-lang/util"
	"log"
)

type User struct {
	ID       int64
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (e *User) Save() error {
	query := `INSERT INTO user (name, email, password) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return err
	}
	defer stmt.Close()
	hashPassword, err := util.HashPassword(e.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	log.Println(hashPassword)
	result, err := stmt.Exec(e.Name, e.Email, hashPassword)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	e.Password = hashPassword
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return err
	}

	e.ID = id
	return nil
}
