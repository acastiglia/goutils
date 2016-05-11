package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var newline rune = '\n'

func main() {
	if len(os.Args) < 2 {
		count(os.Stdin, "")
	} else {
		lines := 0
		words := 0
		chars := 0
		for _, path := range os.Args[1:] {
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			l, w, c := count(file, path)
			lines += l
			words += w
			chars += c
		}

		fmt.Printf("\t%d\t%d\t%d %s\n", lines, words, chars, "total")
	}
}

func count(r io.Reader, path string) (lines int, words int, chars int) {
	sc := bufio.NewScanner(r)
	l := 0
	w := 0
	c := 0
	for sc.Scan() {
		l++
		w += len(strings.Fields(sc.Text()))
		c += len(sc.Text()) + 1
	}

	fmt.Printf("\t%d\t%d\t%d %s\n", l, w, c, path)

	return l, w, c
}
