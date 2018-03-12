package project

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Resolver :
type Resolver struct {
	Base   string
	Strict bool
}

// New :
func New(cwd string, strict bool, pattern string) (*Resolver, error) {
	path, err := lookup(cwd, pattern)
	if err != nil {
		return nil, err
	}
	return &Resolver{Base: path, Strict: strict}, nil
}

// Relative :
func (r *Resolver) Relative(path string) (string, error) {
	result := filepath.Join(r.Base, path)
	if r.Strict {
		if _, err := os.Stat(result); err != nil {
			return "", err
		}
	}
	return result, nil
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
