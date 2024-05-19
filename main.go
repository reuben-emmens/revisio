package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v4"

	"github.com/reuben-emmens/revisio/flashcard"
	"github.com/reuben-emmens/revisio/utils"
)

const (
	currentVersion  = "0.1.0-alpha"
	rootHelpMessage = `Usage: revisio [options]

Options:
  --help 	Print this help message.
  --version	Print version information

Subcommands:
  create	Create a new resource

Add the -h suffix to subcommands for more configuration options`
)

var (
	commandNotFound = errors.New("command not found")
)

func errHandler(ctx utils.Context, err error) {
	fmt.Printf("revisio: %s\n", err)
	ctx.Logr.Println(err)
}

func main() {
	ctx, err := utils.NewContext()
	if err != nil {
		errHandler(ctx, err)
	}

	fs := ff.NewFlagSet("fs")
	var (
		versionFlag = fs.Bool('v', "version", "print version information")
		helpFlag    = fs.Bool('h', "help", "print help documentation")
	)

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("REVISIO"),
	); err != nil {
		errHandler(ctx, err)

	}

	if len(os.Args) < 2 {
		fmt.Println(rootHelpMessage)
		os.Exit(0)
	}

	if *helpFlag {
		fmt.Println(rootHelpMessage)
		os.Exit(0)
	}

	if *versionFlag {
		if err := utils.PrintVersion(currentVersion); err != nil {
			errHandler(ctx, err)
		}
		os.Exit(0)
	}

	switch os.Args[1] {
	case "create":
		if err := flashcard.Create(ctx); err != nil {
			errHandler(ctx, err)
		}
	default:
		errHandler(ctx, commandNotFound)
	}
}
