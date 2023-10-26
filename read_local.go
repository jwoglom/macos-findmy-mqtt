package main

import (
	"fmt"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
)

func main() {
	items, err := fmipcore.ReadItems()
	fmt.Printf("%v %##v\n", err, items)
}
