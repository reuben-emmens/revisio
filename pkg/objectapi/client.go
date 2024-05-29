package objectapi

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Client struct {
	database *database
}

func NewClient() (*Client, error) {
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

func (c *Client) Read(ctx context.Context, key string) ([]string, error) {
	return c.database.read(key)
}

type database struct {
	directory string
	file      string
}

func newDatabase() (*database, error) {
	const subpath = "/.local/share/revisio"

	home, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("error detecting user home directory: %s", err)
		return nil, err
	}

	dirpath := filepath.Join(home, subpath)
	if err = os.MkdirAll(dirpath, 0644); err != nil {
		err = fmt.Errorf("error creating/opening directory, %s: %s", dirpath, err)
		return nil, err
	}

	filepath := filepath.Join(dirpath, "data.csv")
	return &database{
		directory: dirpath,
		file:      filepath,
	}, nil
}

func (db *database) create(key, value string) error {
	file, err := os.OpenFile(db.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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

func (db *database) read(key string) ([]string, error) {
	RecordNotExist := errors.New("no record of this subject has been found")

	file, err := os.OpenFile(db.file, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for record, err := reader.Read(); err == nil; record, err = reader.Read() {
		if record[0] == key {
			return record, nil
		}
	}
	return nil, RecordNotExist
}
