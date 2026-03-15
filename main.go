package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Note struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

const dataFile = "notes.json"

func loadNotes() []Note {
	data, err := os.ReadFile(dataFile)

	if err != nil {
		return []Note{}

	}

	var notes []Note
	json.Unmarshal(data, &notes)
	return notes
}

func saveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func nextID(notes []Note) int {
	max := 0
	for _, n := range notes {
		if n.ID > max {
			max = n.ID
		}
	}
	return max + 1
}

func addNote(title, body string) {
	notes := loadNotes()
	note := Note{
		ID:        nextID(notes),
		Title:     title,
		Body:      body,
		CreatedAt: time.Now().Format("2006-01-02 15:04"),
	}
	notes = append(notes, note)
	if err := saveNotes(notes); err != nil {
		fmt.Println("Error Saving: ", err)
		return
	}
	fmt.Printf("Note %d saved!.\n", note.ID)
}

func listNotes() {
	notes := loadNotes()
	if len(notes) == 0 {
		fmt.Println("Not Notes Yet")
		return
	}

	for _, n := range notes {
		fmt.Printf("#%d [%s] %s\n", n.ID, n.CreatedAt, n.Title)
	}
}

func viewNote(id int) {
	notes := loadNotes()

	for _, n := range notes {
		if n.ID == id {
			fmt.Printf("Title: %s \n Date: %s \n\n %s\n", n.Title, n.CreatedAt, n.Body)
			return

		}
	}
	fmt.Println("Notes not Found.")
}

func deleteNote(id int) {
	notes := loadNotes()
	newNotes := []Note{}
	found := false

	for _, n := range notes {
		if n.ID == id {
			found = true
			continue
		}
		newNotes = append(newNotes, n)

	}
	if !found {
		fmt.Println("Note not Found")
		return
	}
	saveNotes(newNotes)
	fmt.Printf("Note %d Deleted\n", id)

}

func searchNotes(query string) {
	notes := loadNotes()
	query = strings.ToLower(query)
	found := false
	for _, n := range notes {
		if strings.Contains(strings.ToLower(n.Title), query) ||
			strings.Contains(strings.ToLower(n.Body), query) {
			fmt.Printf("#%d [%s] %s\n", n.ID, n.CreatedAt, n.Title)
			found = true
		}

	}
	if !found {
		fmt.Println("No Notes matches")

	}
}

func printUsage() {
	fmt.Println(`Usage:
	note add --title "..." --body "..."
	note list
	note view
	note view --id 1
	note delete --id 1
	note search -- query "..."`)

}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Note Title")
	addBody := addCmd.String("body", "", "Note Body")

	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	viewID := viewCmd.String("id", "", "Note ID")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteID := deleteCmd.String("id", "", "Note ID")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchQuery := searchCmd.String("query", "", "Search Query")

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" {
			fmt.Println("Title is Required")
			return
		}
		addNote(*addTitle, *addBody)
	case "list":
		listNotes()
	case "view":
		viewCmd.Parse(os.Args[2:])
		id, err := strconv.Atoi(*viewID)

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		viewNote(id)

	case "delete":
		deleteCmd.Parse(os.Args[2:])

		id, err := strconv.Atoi(*deleteID)

		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		deleteNote(id)

	case "search":
		searchCmd.Parse(os.Args[2:])

		if *searchQuery == "" {
			fmt.Println("Query is Required")
			return
		}

		searchNotes(*searchQuery)
	default:
		printUsage()

	}

}
