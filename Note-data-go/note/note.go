package note

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	createdAt time.Time `json:"createdAt"`
}

func (note Note) DisplayNote() {
	fmt.Printf("Your Titled Note is %v and content is %v \n", note.Title, note.Content)
}
func (note Note) SaveNote() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json2.Marshal(note)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return os.WriteFile(fileName+".json", json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Pls enter title and content")
	}
	return Note{
		Title:     title,
		Content:   content,
		createdAt: time.Now(),
	}, nil
}
