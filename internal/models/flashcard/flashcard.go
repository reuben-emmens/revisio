package flashcard

import (
	"fmt"
	"io"
)

type Flashcard struct {
	key   string
	value string
}

func (fc *Flashcard) Print(writer io.Writer) {
	fmt.Fprintf(writer, "%s: %s\n", fc.key, fc.value)
}

func (fc *Flashcard) SetKey(s string) {
	fc.key = s
}

func (fc *Flashcard) GetKey() *string {
	return &fc.key
}

func (fc *Flashcard) SetValue(c string) {
	fc.value = c
}

func (fc *Flashcard) GetValue() *string {
	return &fc.value
}
