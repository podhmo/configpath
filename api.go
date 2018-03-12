package configpath

import (
	"github.com/podhmo/configpath/home"
	"github.com/podhmo/configpath/project"
)

type ProjectResolver = project.Resolver
type HomeResolver = home.Resolver

// Project :
func Project(cwd, pattern string) (*ProjectResolver, error) {
	return project.New(cwd, pattern)
}

// Home :
func Home() (*HomeResolver, error) {
	return home.New()
}
