package model

import (
	"errors"
	"leaning-go-lang/db"
	"leaning-go-lang/util"
	"log"
)

type User struct {
	ID       int64
	Name     string
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

func (u *User) ValidateUser() error {
	query := "SELECT password FROM user WHERE email = ? "
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		log.Printf("Error scanning user by email: %v", err)
		return errors.New("InValid credentials")
	}
	log.Println("retrievedPassword: ", retrievedPassword)
	log.Println(" U Password: ", u.Password)
	isValidaPassword := util.DecryptPassword(u.Password, retrievedPassword)
	if !isValidaPassword {
		return errors.New("InValid credentials")
	}
	return nil
}
