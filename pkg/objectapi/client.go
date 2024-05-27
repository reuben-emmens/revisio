package objectapi

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Object is meant to be a domain object for a theoretical object store.
type Object struct {
	Key    string
	Value  string
	Access time.Time
}

type Client struct {
	database        *database
}

func NewClient(token string) (*Client, error) {
	database, err := newDatabase()
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		database: database,
	}, nil
}

func (c *Client) Create(ctx context.Context, key, value string) error {
	return c.database.create(key, value)
}

type database struct {
	location string
}

func newDatabase() (*database, error) {
	const subpath = "/.local/share/revisio"

	home, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("error detecting user home directory: %s", err)
		return &database{}, err
	}

	UserDataDir := filepath.Join(home, subpath)
	if err = os.MkdirAll(UserDataDir, 0644); err != nil {
		err = fmt.Errorf("error accessing user-data directory, %s: %s", UserDataDir, err)
		return &database{}, err
	}

	return &database{
		location:   UserDataDir,
	}, nil
}

func (db *database) create(key, value string) error {
	file, err := os.OpenFile(db.location, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	record := []string{key, value}

	if err := w.Write(record); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return nil
}

// func NewFlashcard(subjectFlag, contentFlag string) (*flashcard, error) {
// 	return &flashcard{subject: subjectFlag, content: contentFlag}, nil
// }



// 	const (
// 		createHelpMessage = `Usage: revisio create [options]
	
// Options:
// 	-s, --subject The subject of the flashcard
// 	-c, --content The content of the flashcard
// 	-h, --help 	  Print this help message.`
// 	)
// 	if len(os.Args) < 3 {
// 		return fmt.Errorf("missing subcommand")
// 	}

