package project

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Resolver :
type Resolver struct {
	Base string
}

// New :
func New(cwd, pattern string) (*Resolver, error) {
	path, err := lookup(cwd, pattern)
	if err != nil {
		return nil, err
	}
	return &Resolver{Base: path}, nil
}

// Relative :
func (r *Resolver) Relative(path string) string {
	return filepath.Join(r.Base, path)
}

func lookup(cwd, pattern string) (string, error) {
	d, err := filepath.Abs(cwd)
	if err != nil {
		return "", err
	}
	sep := string(filepath.Separator)
	elems := strings.Split(d, sep)
	for i := len(elems); i > 0; i-- {
		path := strings.Join(elems[:i], sep)
		if _, err := os.Stat(filepath.Join(path, pattern)); err == nil {
			return path, nil
		}
	}
	return "", io.EOF // xxx :
}
