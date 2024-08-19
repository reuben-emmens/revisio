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

func (c *csvfile) ReadValue(ctx context.Context, key string) (string, error) {
	kV, err := c.read(key)
	if err != nil {
		return "", err
	}

	return kV["value"], nil
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
	kV := make(map[string]string)
	for row, err := reader.Read(); err == nil; row, err = reader.Read() {
		if row[0] == key {
			for i, value := range row {
				switch i {
				case 0:
					kV["key"] = key
				case 1:
					kV["value"] = value
				}
			}
			return kV, nil
		}
	}
	return nil, ErrNotExist
}
