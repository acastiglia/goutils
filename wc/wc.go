package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		count(os.Stdin)
	} else {
		lines, words, chars := 0, 0, 0
		for _, path := range os.Args[1:] {
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			l, w, c, _ := count(file)
			lines += l
			words += w
			chars += c
			fmt.Printf("\t%d\t%d\t%d %s\n", lines, words, chars, path)
		}

		if len(os.Args) > 2 {
			fmt.Printf("\t%d\t%d\t%d %s\n", lines, words, chars, "total")
		}
	}
}

func count(reader io.Reader) (lines int, words int, chars int, bytes int) {
	l, w, c, b := 0, 0, 0, 0
	runeReader := bufio.NewReader(reader)

	for {
		lw, lc, lb, err := readLine(runeReader)
		if err == io.EOF {
			return l, w, c, b
		}
		if err != nil {
			panic(err)
		}
		l++
		w += lw
		c += lc
		b += lb
	}

	return l, w, c, b
}

func readLine(reader *bufio.Reader) (int, int, int, error) {
	w, c, b := 0, 0, 0
	prevRune := ' '
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			return w, c, b, err
		}
		c++
		b += size
		if !unicode.IsSpace(r) && unicode.IsSpace(prevRune) {
			w++
		}
		if r == '\n' {
			return w, c, b, nil
		}
		prevRune = r
	}

	return w, c, b, nil
}
