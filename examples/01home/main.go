package main

import (
	"fmt"

	"github.com/podhmo/configpath"
)

func main() {
	r, _ := configpath.Home()
	fmt.Println(configpath.Must(r.Relative(".config")))
}
