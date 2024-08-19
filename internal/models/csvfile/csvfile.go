package csvfile

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrNotExist = errors.New("no flashcard with this subject exists")
)

type csvfile struct {
	directory string
	file      string
}

func New(ctx context.Context) (*csvfile, error) {
	csvFilePtr, err := newCsvfile()
	if err != nil {
		return nil, err
	}
	return csvFilePtr, nil
}

func (c *csvfile) Create(ctx context.Context, key, value string) error {
	return c.create(key, value)
}

func (c *csvfile) Read(ctx context.Context, key string) (map[string]string, error) {
	qAndA, err := c.read(key)
	if err != nil {
		return nil, err
	}

	return qAndA, nil
}

func newCsvfile() (*csvfile, error) {
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
	return &csvfile{
		directory: dirpath,
		file:      filepath,
	}, nil
}

func (c *csvfile) create(key, value string) error {
	file, err := os.OpenFile(c.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
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

func (c *csvfile) read(key string) (map[string]string, error) {

	file, err := os.OpenFile(c.file, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for line, err := reader.Read(); err == nil; line, err = reader.Read() {
		if line[0] == key {
			value := line[1]
			qAndA := make(map[string]string)
			qAndA["question"] = key
			qAndA["answer"] = value

			return qAndA, nil
		}
	}
	return nil, ErrNotExist
}
