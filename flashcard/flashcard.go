package flashcard

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/v4"
	"github.com/reuben-emmens/revisio/utils"
)

type flashcard struct {
	subject string
	content string
}

func NewFlashcard(subjectFlag, contentFlag string) (*flashcard, error) {
	return &flashcard{subject: subjectFlag, content: contentFlag}, nil
}

func WriteToCsv(fc *flashcard, ctx utils.Context) error {
	file, err := os.OpenFile(ctx.Data, os.O_WRONLY|os.O_CREATE, 0644)
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

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func Create(ctx utils.Context) error {

	const (
		createHelpMessage = `Usage: revisio create [options]
	
Options:
	-s, --subject The subject of the flashcard
	-c, --content The content of the flashcard
	-h, --help 	  Print this help message.`
	)
	if len(os.Args) < 3 {
		return fmt.Errorf("missing subcommand")
	}

	switch os.Args[2] {
	case "flashcard":
		createFs := ff.NewFlagSet("createFs")

		var (
			subjectFlag = createFs.String('s', "subject", "", "subject of the flashcard")
			contentFlag = createFs.String('c', "content", "", "content of the flashcard")
			helpFlag    = createFs.Bool('h', "help", "print help documentation")
		)

		if err := ff.Parse(createFs, os.Args[3:]); err != nil {
			log.Fatalf("error parsing create variables: %s", err)
		}
		if len(os.Args) < 4 {
			fmt.Println(createHelpMessage)
			os.Exit(0)
		}

		if *helpFlag {
			fmt.Println(createHelpMessage)
			os.Exit(0)
		}

		fc, err := NewFlashcard(*subjectFlag, *contentFlag)
		if err != nil {
			return err
		}

		if err := WriteToCsv(fc, ctx); err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid create subcommand")
	}
	return nil
}
