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

	_, err = exec.LookPath(prog)
	if err != nil {
		log.Fatalf("please, install %s first.", prog)
	}

	// cmd := exec.Command(prog, "build", "-o test", "-gcflags -m", file.Name())
	cmd := exec.Command(prog, "-version")

	cmd.Env = []string{
		fmt.Sprintf("GOBIN=%s", os.Getenv("GOBIN")),
		fmt.Sprintf("GOPATH=%s", os.Getenv("GOPATH")),
		fmt.Sprintf("GOROOT=%s", os.Getenv("GOROOT")),
		"GOARCH=amd64", "GOHOSTARCH=amd64", "GOHOSTOS=linux", "GOOS=linux"}

	fmt.Printf("%v\n", cmd.Env)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("out: %s\n", out)

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
