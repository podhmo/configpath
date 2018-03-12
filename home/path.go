package home

import (
	"os/user"
	"path/filepath"
)

// Resolver :
type Resolver struct {
	U *user.User
}

// New :
func New() (*Resolver, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &Resolver{U: u}, nil
}

// Relative :
func (r *Resolver) Relative(name string) string {
	return filepath.Join(r.U.HomeDir, name)
}

// ConfigDir :
func (r *Resolver) ConfigDir(name string) string {
	return r.Relative(filepath.Join(".config", name))
}
