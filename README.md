# configpath

## project

```go
package main

import (
	"fmt"

	"github.com/podhmo/configpath"
)

func main() {
	// lookup the directory has "bin" directory.
	r, _ := configpath.Project("bin")

	fmt.Println(configpath.Must(r.Relative("bin/goimports")))
	// e.g. /home/podhmo/go/bin/goimports
}
```

## home

```go
package main

import (
	"fmt"

	"github.com/podhmo/configpath"
)

func main() {
	r, _ := configpath.Home()

	fmt.Println(configpath.Must(r.Relative(".config")))
	// e.g. /home/podhmo/.config
}
```
