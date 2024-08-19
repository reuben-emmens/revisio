package createcmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffval"

	"github.com/reuben-emmens/revisio/pkg/rootcmd"
)

type CreateConfig struct {
	*rootcmd.RootConfig
	Subject string
	Content string
	Flags   *ff.FlagSet
	Command *ff.Command
}

func New(rootConfig *rootcmd.RootConfig) *CreateConfig {
	var cfg CreateConfig
	cfg.RootConfig = rootConfig
	cfg.Flags = ff.NewFlagSet("create").SetParent(cfg.RootConfig.Flags)
	_, err := cfg.Flags.AddFlag(ff.FlagConfig{
		ShortName: 'k',
		LongName:  "key",
		Value:     ffval.NewValue(&cfg.Subject),
		Usage:     "The key of the flashcard",
		NoDefault: true,
	})
	if err != nil {
		fmt.Fprintln(cfg.Stderr, rootcmd.AddFlagErr.Error())
	}
	_, err = cfg.Flags.AddFlag(ff.FlagConfig{
		ShortName: 'v',
		LongName:  "value",
		Value:     ffval.NewValue(&cfg.Content),
		Usage:     "The value of the flashcard",
		NoDefault: true,
	})
	if err != nil {
		fmt.Fprintln(cfg.Stderr, rootcmd.AddFlagErr.Error())
	}
	cfg.Command = &ff.Command{
		Name:      "create",
		Usage:     "revisio create [FLAGS] <KEY> <VALUE>",
		ShortHelp: "create a flashcard",
		Flags:     cfg.Flags,
		Exec:      cfg.Exec,
	}
	cfg.RootConfig.Command.Subcommands = append(cfg.RootConfig.Command.Subcommands, cfg.Command)
	return &cfg
}

func (cfg *CreateConfig) Exec(ctx context.Context, args []string) error {
	if cfg.Subject == "" {
		return errors.New("Missing required flag: -k, --key")
	} else if cfg.Content == "" {
		return errors.New("Missing required flag: -v, --value")
	}

	if err := cfg.Client.Create(ctx, cfg.Subject, cfg.Content); err != nil {
		return err
	}

	if cfg.Verbose {
		fmt.Fprintf(cfg.Stderr, "create %q OK\n", cfg.Subject)
	}

	return nil
}
