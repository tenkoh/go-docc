package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/tenkoh/go-docc"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		panic("must input a MS-Word filename as an argument")
	}
	r, err := docc.NewReader(args[0])
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for {
		p, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(p)
	}
}
