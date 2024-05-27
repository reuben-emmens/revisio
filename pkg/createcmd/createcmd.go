package createcmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/peterbourgon/ff/v4"
	"github.com/reuben-emmens/revisio/pkg/rootcmd"
)

type CreateConfig struct {
	*rootcmd.RootConfig
	Flags     *ff.FlagSet
	Command   *ff.Command
}

func New(rootConfig *rootcmd.RootConfig) *CreateConfig {
	var cfg CreateConfig
	cfg.RootConfig = rootConfig
	cfg.Flags = ff.NewFlagSet("create").SetParent(cfg.RootConfig.Flags)
	cfg.Command = &ff.Command{
		Name:      "create",
		Usage:     "objectctl create [FLAGS] <KEY> <VALUE>",
		ShortHelp: "create a flashcard",
		Flags:     cfg.Flags,
		Exec:      cfg.Exec,
	}
	cfg.RootConfig.Command.Subcommands = append(cfg.RootConfig.Command.Subcommands, cfg.Command)
	return &cfg
}

func (cfg *CreateConfig) Exec(ctx context.Context, args []string) error {
	if len(args) < 2 {
		return errors.New("create requires at least 2 args")
	}

	var (
		key   = args[0]
		value = strings.Join(args[1:], " ")
		err   = cfg.Client.Create(ctx, key, value)
	)
	if err != nil {
		return err
	}

	if cfg.Verbose {
		fmt.Fprintf(cfg.Stderr, "create %q OK\n", key)
	}

	return nil
}
