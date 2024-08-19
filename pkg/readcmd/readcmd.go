package readcmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffval"

	"github.com/reuben-emmens/revisio/pkg/rootcmd"
)

type ReadConfig struct {
	*rootcmd.RootConfig
	Flags   *ff.FlagSet
	Command *ff.Command
}

func New(rootConfig *rootcmd.RootConfig) *ReadConfig {
	var cfg ReadConfig
	cfg.RootConfig = rootConfig
	cfg.Flags = ff.NewFlagSet("read").SetParent(cfg.RootConfig.Flags)
	_, err := cfg.Flags.AddFlag(ff.FlagConfig{
		ShortName: 'k',
		LongName:  "key",
		Value:     ffval.NewValue(cfg.Flashcard.GetKey()),
		Usage:     "The key of the flashcard",
		NoDefault: true,
	})
	if err != nil {
		fmt.Fprintln(cfg.Stderr, rootcmd.AddFlagErr.Error())
	}
	cfg.Command = &ff.Command{
		Name:      "read",
		Usage:     "revisio read [FLAGS] <KEY> <VALUE>",
		ShortHelp: "read a flashcard",
		Flags:     cfg.Flags,
		Exec:      cfg.Exec,
	}
	cfg.RootConfig.Command.Subcommands = append(cfg.RootConfig.Command.Subcommands, cfg.Command)
	return &cfg
}

func (cfg *ReadConfig) Exec(ctx context.Context, args []string) error {
	subject := cfg.Flashcard.GetKey()
	if *subject == "" {
		return errors.New("Missing required flag: -k, --key")
	}

	value, err := cfg.Client.ReadValue(ctx, *subject)
	if err != nil {
		return err
	}

	cfg.Flashcard.SetValue(value)

	if cfg.Verbose {
		fmt.Fprintf(cfg.Stderr, "read %q OK\n", *subject)
	}

	cfg.Flashcard.Print(cfg.Stderr)

	return nil
}
