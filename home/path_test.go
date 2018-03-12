package home

import (
	"fmt"
	"os/user"
	"path/filepath"
	"testing"
)

// TestRelative :
func TestRelative(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	for i, c := range []struct {
		path     string
		expected string
	}{
		{
			path:     "",
			expected: fmt.Sprintf("%s", u.HomeDir),
		},
		{
			path:     "foo",
			expected: fmt.Sprintf("%s/foo", u.HomeDir),
		},
		{
			path:     ".foo",
			expected: fmt.Sprintf("%s/.foo", u.HomeDir),
		},
		{
			path: "../foo",
			expected: func() string {
				path, _ := filepath.Abs(fmt.Sprintf("%s/../foo", u.HomeDir))
				return path
			}(),
		},
	} {
		c := c
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got, _ := Relative(c.path)
			if c.expected != got {
				t.Errorf("expected %q, but got %q", c.expected, got)
			}
		})
	}

}
