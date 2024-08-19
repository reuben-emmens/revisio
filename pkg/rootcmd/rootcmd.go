package rootcmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffval"
	"github.com/reuben-emmens/revisio/internal/models/client"
	"github.com/reuben-emmens/revisio/internal/models/flashcard"
)

var (
	AddFlagErr = errors.New("unable to add flag")
)

type RootConfig struct {
	Stdout    io.Writer
	Stderr    io.Writer
	Verbose   bool
	Flags     *ff.FlagSet
	Command   *ff.Command
	Client    client.Client
	Flashcard *flashcard.Flashcard
}

func New(stdout, stderr io.Writer) *RootConfig {
	var cfg RootConfig
	cfg.Stdout = stdout
	cfg.Stderr = stderr
	cfg.Flags = ff.NewFlagSet("revisio")
	_, err := cfg.Flags.AddFlag(ff.FlagConfig{
		ShortName: 'v',
		LongName:  "verbose",
		Value:     ffval.NewValue(&cfg.Verbose),
		Usage:     "log verbose output",
		NoDefault: true,
	})
	if err != nil {
		fmt.Fprintln(cfg.Stderr, AddFlagErr.Error())
	}
	cfg.Command = &ff.Command{
		Name:      "revisio",
		ShortHelp: "manage flashcards",
		Usage:     "revisio [FLAGS] <SUBCOMMAND> ...",
		Flags:     cfg.Flags,
	}

	cfg.Flashcard = new(flashcard.Flashcard)

	return &cfg
}
