// package createcmd_test

// import (
// 	"testing"

// 	"github.com/reuben-emmens/revisio/pkg/create"
// )

// func TestNewFlashCard(t *testing.T) {
// 	scenarios := []struct {
// 		subject, content string
// 	}{
// 		{subject: "revisio", content: "a flashcard app"},
// 	}

// 	for _, s := range scenarios {
// 		fc, err := flashcard.NewFlashcard(s.subject, s.content)
// 		if err != nil {
// 			t.Errorf("error creating new flashcard: %s", err)
// 		}

// 		t.Logf("%+v", fc)
// 	}
// }