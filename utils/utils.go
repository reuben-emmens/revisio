package utils

import (
	"fmt"

	versionPkg "github.com/hashicorp/go-version"
)

func PrintVersion(version string) error {
	newVersion, err := versionPkg.NewVersion(version)
	if err != nil {
		return err
	}

	fmt.Printf("revisio version %s\n", newVersion.String())
	return nil
}
