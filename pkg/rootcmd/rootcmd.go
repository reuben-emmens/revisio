package rootcmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffval"
	"github.com/reuben-emmens/revisio/pkg/objectapi"
)

var (
	AddFlagErr = errors.New("unable to add flag")
)

type RootConfig struct {
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
	Client  *objectapi.Client
	Flags   *ff.FlagSet
	Command *ff.Command
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
	return &cfg
}
