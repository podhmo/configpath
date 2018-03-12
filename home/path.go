package home

import (
	"os"
	"os/user"
	"path/filepath"
)

// Resolver :
type Resolver struct {
	U      *user.User
	Strict bool
}

// New :
func New(strict bool) (*Resolver, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &Resolver{U: u, Strict: strict}, nil
}

// Relative :
func (r *Resolver) Relative(name string) (string, error) {
	result := filepath.Join(r.U.HomeDir, name)
	if r.Strict {
		if _, err := os.Stat(result); err != nil {
			return "", err
		}
	}
	return result, nil
}

// ConfigDir :
func (r *Resolver) ConfigDir(name string) (string, error) {
	return r.Relative(filepath.Join(".config", name))
}
