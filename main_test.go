package main

import (
	"testing"

	repro "github.com/qsq8t65pxs/dep-go-repro-20260513"
)

func TestDependencyMessage(t *testing.T) {
	if repro.Message() == "" {
		t.Fatal("empty message")
	}
}
