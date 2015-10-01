package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var output string

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-out=out.path] in.path\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&output, "out", "out.go", "Specify a path to the output file")

	flag.Parse()
}

func main() {
	checkRequirements()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer file.Close()

	prog := "go"
	path, err := exec.LookPath(prog)
	if err != nil {
		log.Fatalf("please, install %s first.", prog)
	}
	fmt.Printf("%s is available at %s\n", prog, path)

	fmt.Printf("input file: %s, output file: %s\n", flag.Arg(0), output)
}

func checkRequirements() {
	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()

		log.Fatalf("the input file not specified\n")
	} else if len(args) > 1 {
		log.Printf("to many positional arguments, ignoring %v\n", args[1:])
	}
}
