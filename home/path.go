package home

import (
	"os/user"
	"path/filepath"
)

// Relative :
func Relative(name string) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, name), nil
}

// ConfigDir :
func ConfigDir(name string) (string, error) {
	return Relative(filepath.Join(".config", name))
}
