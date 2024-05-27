package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"

	"github.com/reuben-emmens/revisio/pkg/createcmd"
	"github.com/reuben-emmens/revisio/pkg/objectapi"
	"github.com/reuben-emmens/revisio/pkg/rootcmd"
	"github.com/reuben-emmens/revisio/pkg/versioncmd"
)

func main() {
	var (
		ctx    = context.Background()
		args   = os.Args[1:]
		stdout = os.Stdout
		stderr = os.Stderr
		err    = exec(ctx, args, stdout, stderr)
	)
	switch {
	case err == nil, errors.Is(err, ff.ErrHelp), errors.Is(err, ff.ErrNoExec):
		// no problem
	case err != nil:
		fmt.Fprintf(stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func exec(ctx context.Context, args []string, stdout, stderr io.Writer) (err error) {
	var (
		root = rootcmd.New(stdout, stderr)
		_    = versioncmd.New(root)
		_    = createcmd.New(root)
	)

	defer func() {
		if err != nil {
			fmt.Fprintf(stderr, "\n%s\n", ffhelp.Command(root.Command))
		}
	}()

	if err := root.Command.Parse(args); err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	client, err := objectapi.NewClient()
	if err != nil {
		return fmt.Errorf("construct API client: %w", err)
	}

	root.Client = client

	if err := root.Command.Run(ctx); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}
