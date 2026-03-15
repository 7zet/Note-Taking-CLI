# 📝 Note-Taking CLI

A simple command-line note manager written in Go.
Notes are stored locally in a `notes.json` file.

## Usage
```bash
# Add a note
go run main.go add --title "Meeting" --body "Discuss project goals"

# List all notes
go run main.go list

# View a note
go run main.go view --id 1

# Search notes
go run main.go search --query "meeting"

# Delete a note
go run main.go delete --id 1
```

## What I learned

- CLI argument parsing with the `flag` package
- Reading and writing JSON files with `encoding/json`
- Structs, slices, and functions in Go

## Built with

- Go 1.22+
- Standard library only
