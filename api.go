package configpath

import (
	"github.com/podhmo/configpath/home"
	"github.com/podhmo/configpath/project"
)

type ProjectResolver = project.Resolver
type HomeResolver = home.Resolver

// Project :
func Project(pattern string) (*ProjectResolver, error) {
	cwd := ""
	strict := true
	return project.New(cwd, strict, pattern)
}

// Home :
func Home() (*HomeResolver, error) {
	strict := true
	return home.New(strict)
}

// Must :
func Must(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}
