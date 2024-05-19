package flashcard

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/peterbourgon/ff/v4"
)

type flashcard struct {
	subject string
	content string
}

func NewFlashcard(subjectFlag, contentFlag string) (*flashcard, error) {
	return &flashcard{subject: subjectFlag, content: contentFlag}, nil
}

func WriteToCsv(fc *flashcard, csvPath string) error {
	if err := os.MkdirAll(filepath.Dir(csvPath), 0744); err != nil {
		return err
	}

	file, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	record := make([]string, 2)
	record[0] = fc.subject
	record[1] = fc.content

	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func GetDataPath(path string) (string, error) {
	if path != "/.local/share/revisio/data.csv" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, path), nil
}

func Create(csvPath string) error {
	if len(os.Args) < 3 {
		return fmt.Errorf("Missing create subcommand")
	}
	switch os.Args[2] {
	case "flashcard":
		createFs := ff.NewFlagSet("new")

		var (
			subjectFlag = createFs.String('s', "subject", "Hello World!", "subject of the flashcard")
			contentFlag = createFs.String('c', "content", "A standard 'first-project", "content of the flashcard")
		)

		if err := ff.Parse(createFs, os.Args[3:]); err != nil {
			log.Fatalf("error parsing create variables: %s", err)
		}

		fc, err := NewFlashcard(*subjectFlag, *contentFlag)
		if err != nil {
			return err
		}

		if err := WriteToCsv(fc, csvPath); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Invalid create subcommand")
	}
	return nil
}
