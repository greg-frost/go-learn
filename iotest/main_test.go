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
