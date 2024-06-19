package main

import (
	"io"
	"testing"
)

func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}
