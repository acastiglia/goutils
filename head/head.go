package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func tryOpen(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s: could not open file\n", path))
	}

	return file, err
}

func head(file *os.File, count int) {
	reader := bufio.NewReader(file)
	for i := 0; i < count; i++ {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}

func main() {
	var count = flag.Int("n", -1, "The number of lines to print")
	flag.Parse()

	if *count < 1 {
		os.Stderr.WriteString(fmt.Sprintf("illegal line count -- %d\n", *count))
		os.Exit(1)
	}

	if len(flag.Args()) == 1 {
		file, err := tryOpen(flag.Args()[0])
		if err != nil {
			os.Exit(1)
		}
		head(file, *count)
		os.Exit(0)
	}

	for i := 0; i < len(flag.Args()); i++ {
		path := flag.Args()[i]
		file, err := tryOpen(path)
		if err != nil {
			continue
		}

		fmt.Println("==>", path)
		head(file, *count)
		if i < len(flag.Args())-1 {
			fmt.Print("\n")
		}
	}
}
