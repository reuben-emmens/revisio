package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	versionPkg "github.com/hashicorp/go-version"
	"github.com/peterbourgon/ff"
)

const (
	currentVersion  string = "0.1.0-alpha"
	rootHelpMessage string = `Usage: revisio [options]

Options:
  --help 	Print this help message.
  --version	Print version information

Subcommands:
  create	Create a new resource

Add the -h suffix to subcommands for more configuration options`
)

type flashCard struct {
	subject string
	content string
}

func newFlashcard(subjectFlag, contentFlag string) (*flashCard, error) {
	return &flashCard{subject: subjectFlag, content: contentFlag}, nil
}

func printVersion(version string) error {
	newVersion, err := versionPkg.NewVersion(version)
	if err != nil {
		return err
	}

	fmt.Printf("revisio version %s\n", newVersion.String())
	return nil
}

func writeToCsv(fc *flashCard, csvPath string) error {
	if err := os.MkdirAll(filepath.Dir(csvPath), 0644); err != nil {
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

func getDataPath(path string) (string, error) {
	if path != "/.local/share/revisio/data.csv" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, path), nil
}

func create(csvPath string) error {
	if len(os.Args) < 3 {
		return fmt.Errorf("Missing create subcommand")
	}
	switch os.Args[2] {
	case "flashcard":
		createFs := flag.NewFlagSet("new", flag.ExitOnError)

		var (
			subjectFlag = createFs.String("subject", "Hello World!", "subject of the flashcard")
			contentFlag = createFs.String("content", "A standard 'first-project", "content of the flashcard")
		)

		if err := ff.Parse(createFs, os.Args[3:]); err != nil {
			log.Fatalf("error parsing create variables: %s", err)
		}

		fc, err := newFlashcard(*subjectFlag, *contentFlag)
		if err != nil {
			return err
		}

		if err := writeToCsv(fc, csvPath); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Invalid create subcommand")
	}
	return nil
}

func main() {
	fs := flag.NewFlagSet("fs", flag.ExitOnError)
	var (
		versionFlag = fs.Bool("version", false, "print version information")
		helpFlag    = fs.Bool("help", false, "print help documentation")
		csvPathFlag = fs.String("csv-path", "/.local/share/revisio/data.csv", "location of .csv file containing cards")
	)

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("REVISIO"),
	); err != nil {
		log.Fatalf("error parsing variables: %s", err)
	}

	if len(os.Args) < 2 {
		log.Println("Missing subcommand")
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Println(rootHelpMessage)
		os.Exit(0)
	}

	if *versionFlag {
		if err := printVersion(currentVersion); err != nil {
			log.Fatalf("%s", err)
		}
		return
	}

	csvPath, err := getDataPath(*csvPathFlag)
	if err != nil {
		log.Fatalf("error detecting user home directory: %s", err)
	}

	switch os.Args[1] {
	case "create":
		if err := create(csvPath); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Invalid subcommand")
	}

}
