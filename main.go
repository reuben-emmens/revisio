package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/v4"

	"github.com/reuben-emmens/revisio/flashcard"
	"github.com/reuben-emmens/revisio/utils"
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

func main() {
	fs := ff.NewFlagSet("fs")
	var (
		versionFlag = fs.Bool('v', "version", "print version information")
		helpFlag    = fs.Bool('h', "help", "print help documentation")
		csvPathFlag = fs.String('p', "path", "/.local/share/revisio/data.csv", "path to .csv file to store flashcard data")
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
		if err := utils.PrintVersion(currentVersion); err != nil {
			log.Fatalf("%s", err)
		}
		return
	}

	csvPath, err := flashcard.GetDataPath(*csvPathFlag)
	if err != nil {
		log.Fatalf("error detecting user home directory: %s", err)
	}

	switch os.Args[1] {
	case "create":
		if err := flashcard.Create(csvPath); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Invalid subcommand")
	}
}
