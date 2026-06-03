package main

import (
	"strings"
	"testing"
	"testing/iotest"
)

func TestLowerCaseReader(t *testing.T) {
	err := iotest.TestReader(
		&LowerCaseReader{
			reader: strings.NewReader("HHHheLlLlOoWworRRldD"),
		},
		[]byte("helloworld"),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPrint(t *testing.T) {
	err := Print(iotest.TimeoutReader(
		strings.NewReader("Timeout reader"),
	))
	if err != nil {
		t.Fatal(err)
	}
}
