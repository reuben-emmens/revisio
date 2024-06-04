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
	Subject string
	Flags   *ff.FlagSet
	Command *ff.Command
}

func New(rootConfig *rootcmd.RootConfig) *ReadConfig {
	var cfg ReadConfig
	cfg.RootConfig = rootConfig
	cfg.Flags = ff.NewFlagSet("read").SetParent(cfg.RootConfig.Flags)
	_, err := cfg.Flags.AddFlag(ff.FlagConfig{
		ShortName: 's',
		LongName:  "subject",
		Value:     ffval.NewValue(&cfg.Subject),
		Usage:     "The subject of the flashcard",
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
	if cfg.Subject == "" {
		return errors.New("Missing required flag: -s, --subject")
	}

	record, err := cfg.Client.Read(ctx, cfg.Subject)
	if err != nil {
		return err
	}

	if cfg.Verbose {
		fmt.Fprintf(cfg.Stderr, "read %q OK\n", cfg.Subject)
	}

	fmt.Fprintf(cfg.Stderr, "%s: %s\n", record[0], record[1])

	return nil
}
