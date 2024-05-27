package utils_test

import (
	"testing"

	"github.com/reuben-emmens/revisio/pkg/utils"
)

func TestPrintVersion(t *testing.T) {
	scenarios := []struct {
		version string
	}{
		{version: "0.1.0-alpha"},
	}

	for _, s := range scenarios {
		if err := utils.PrintVersion(s.version); err != nil {
			t.Errorf("error parsing and printing version string: %s", err)
		}
	}
}
