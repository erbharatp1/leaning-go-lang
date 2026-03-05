package main

import (
	"Note-data-go/note"
	"bufio"
	"fmt"
	"os"
)

type saver interface {
	SaveNote() error
}

func getNoteData() (string, string, error) {
	title, err := getUserInfo("Notes title:")
	if err != nil {

		return "", "", err
	}
	content, err := getUserInfo("Notes content:")
	if err != nil {
		return "", "", err
	}
	return title, content, nil
}

func main() {
	title, content, err := getNoteData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	note.DisplayNote()

	err = saveData(note)

	//err = note.SaveNote()
	//if err != nil {
	//	fmt.Println("Error during saving the note:", err)
	//	return
	//}

	//	fmt.Println("Title:", title)
	//fmt.Println("Content:", content)
}
func saveData(data saver) error {
	err := data.SaveNote()
	if err != nil {
		fmt.Println("Error during saving the note:", err)
		return err
	}
	return nil
}

func getUserInfo(promptTest string) (string, error) {
	fmt.Print("\n", promptTest)
	//var input string
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	//text = strings.TrimSpace(reader, "\r")
	if err != nil {
		return "", err
	}
	//reader = strings.TrimSpace(reader, "\n")
	//reader = strings.TrimSpace(reader, "\r")
	return text, nil
}
