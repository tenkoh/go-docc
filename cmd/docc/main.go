package main

import (
	"flag"
	"fmt"

	"github.com/tenkoh/docc"
)

var Version = "v0.0.1"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("must input a MS-Word filename as an argument")
	}
	ps, err := docc.Decode(args[0])
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		fmt.Println(p)
	}
}