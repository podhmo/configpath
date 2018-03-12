package project

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	r, err := New("", false, ".git")
	if err != nil {
		t.Fatal(err)
	}
	got, err := r.Relative("LICENSE")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(got); err != nil {
		t.Fatalf("relative README.md. but got %s", err)
	}
}
