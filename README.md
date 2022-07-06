# go-battlescribe

A Go module for interfacing with various aspects of Battlescribe data.

## Install

`go get github.com/myminicommission/go-battlescribe`

## Usage

### BSData

#### Getting BSData Module

```golang
package main

import (
	"fmt"
	"strings"

	"github.com/myminicommission/go-battlescribe"
)

const (
  wh40k = "wh40k"
)

func main() {
  battlescribe.GetBSData(wh40k)
}
```

### Rosters

#### Read `.rosz` File

```golang
package main

import (
	"fmt"
	"strings"

	"github.com/myminicommission/go-battlescribe"
)

const (
	rosterFilePath = "./Double Tantalus.rosz"
)

func main() {

	r, err := battlescribe.ReadRoszFile(rosterFilePath)
	if err != nil {
		panic(err)
	}

	println(r.Name)
	println(r.GameSystemName)
	for _, cost := range r.Costs.Cost {
		println(fmt.Sprintf("%s - %s", strings.TrimSpace(cost.Name), cost.Value))
	}
}
```