package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const maxPadding = 5

func padding(n int) string {
	return strings.Repeat(" ", maxPadding-int(math.Log10(float64(n))))
}

func nl(file *os.File) {
	reader := bufio.NewReader(file)
	i := 1
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) > 0 {
			fmt.Printf("%s%d\t%s\n", padding(i), i, line)
			i++
		} else {
			fmt.Println()
		}
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		nl(os.Stdin)
		os.Exit(0)
	}

	for _, path := range flag.Args() {
		file, err := os.Open(path)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%s: could not open file\n", path))
			continue
		}

		nl(file)
	}
}
