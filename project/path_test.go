package project

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	r, err := New("", ".git")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(r.Relative("LICENSE")); err != nil {
		t.Fatalf("relative README.md. but got %s", err)
	}
}
