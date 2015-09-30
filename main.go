package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s -out=out.path<optional> in.path<required>\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	checkRequirements()

	fmt.Printf("%v", flag.Args())
}

func checkRequirements() {
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()

		fmt.Printf("Warning: The input file is required\n")

		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("Notice: to many positional arguments, ignoring %v\n", args[1:])
	}
}
