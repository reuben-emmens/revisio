package versioncmd

import (
	"context"
	"fmt"

	"github.com/peterbourgon/ff/v4"

	versionPkg "github.com/hashicorp/go-version"
	"github.com/reuben-emmens/revisio/pkg/rootcmd"
)

const version = "0.1.0-alpha"

type VersionConfig struct {
	*rootcmd.RootConfig
	Version *versionPkg.Version
	Flags   *ff.FlagSet
	Command *ff.Command
}

func New(rootConfig *rootcmd.RootConfig) *VersionConfig {
	var cfg VersionConfig
	cfg.RootConfig = rootConfig
	cfg.Version, _ = versionPkg.NewVersion(version)
	cfg.Flags = ff.NewFlagSet("version").SetParent(cfg.RootConfig.Flags)
	cfg.Command = &ff.Command{
		Name:      "version",
		Usage:     "revisio version",
		ShortHelp: "print version information",
		Flags:     cfg.Flags,
		Exec:      cfg.Exec,
	}
	cfg.RootConfig.Command.Subcommands = append(cfg.RootConfig.Command.Subcommands, cfg.Command)
	return &cfg
}

func (cfg *VersionConfig) Exec(ctx context.Context, args []string) error {
	fmt.Printf("revisio version %s\n", cfg.Version.String())
	return nil
}
