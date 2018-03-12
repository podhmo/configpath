package main

import (
	"fmt"

	"github.com/podhmo/configpath"
)

func main() {
	r, _ := configpath.Project("bin")
	fmt.Println(configpath.Must(r.Relative("bin/goimports")))
}
