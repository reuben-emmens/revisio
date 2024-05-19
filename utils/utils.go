package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	versionPkg "github.com/hashicorp/go-version"
)

func newLogger(userData string) (*log.Logger, error) {
	logPath := filepath.Join(userData, "log.txt")
	file, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}

	logr := log.New(file, "", log.LstdFlags)
	return logr, nil
}

type Context struct {
	UserDataDir string
	Data        string
	Logr        *log.Logger
}

func NewContext() (ctx Context, err error) {
	const subpath = "/.local/share/revisio"
	ctx = Context{}

	home, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("error detecting user home directory: %s", err)
		return
	}

	ctx.UserDataDir = filepath.Join(home, subpath)
	if err = os.MkdirAll(ctx.UserDataDir, 0644); err != nil {
		err = fmt.Errorf("error accessing user-data directory, %s: %s", ctx.UserDataDir, err)
		return
	}
	ctx.Data = filepath.Join(home, subpath, "data.csv")
	ctx.Logr, err = newLogger(ctx.UserDataDir)
	if err != nil {
		return
	}
	return
}

func PrintVersion(version string) error {
	newVersion, err := versionPkg.NewVersion(version)
	if err != nil {
		return err
	}

	fmt.Printf("revisio version %s\n", newVersion.String())
	return nil
}
