package main

import (
	"testing"
)

func TestPrintVersion(t *testing.T) {
	scenarios := []struct {
		version string
	}{
		{version: "0.1.0-alpha"},
	}

	for _, s := range scenarios {
		if err := printVersion(s.version); err != nil {
			t.Errorf("error parsing and printing version string: %s", err)
		}
	}
}

func TestNewFlashCard(t *testing.T) {
	scenarios := []struct {
		subject, content string
	}{
		{subject: "revisio", content: "a flashcard app"},
	}

	for _, s := range scenarios {
		fc, err := newFlashcard(s.subject, s.content)
		if err != nil {
			t.Errorf("error creating new flashcard: %s", err)
		}

		t.Logf("%+v", fc)
	}
}
