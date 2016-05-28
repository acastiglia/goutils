package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	omitNewline := flag.Bool("n", false, "Do not print a trailing newline character")
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), " "))
	if !*omitNewline {
		fmt.Println()
	}
}
