package main

import (
	"bufio"
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

	cmd := exec.Command(prog, "build", "-o", "test", "-gcflags", "-m", file.Name())

	out, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Command execution failed: %s", err)
	}

	line := bufio.NewScanner(out)
	defer out.Close()

	for line.Scan() {
		fmt.Printf("line: %s\n", line.Text())
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Command execution failed: %s", err)
	}

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
