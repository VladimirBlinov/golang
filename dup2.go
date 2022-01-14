package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	cntfiles := make(map[string]int)
	files := os.Args
	if len(files) == 0 {
		fmt.Println("no files")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %v\n", err)
				continue
			}
			countLines(f, counts, cntfiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	for filename, fn := range cntfiles {
		fmt.Printf("%d\t%s\n", fn, filename)
	}

}

func countLines(f *os.File, counts map[string]int, cntfiles map[string]int) {
	input := bufio.NewScanner(f)
	cnt := 0
	for input.Scan() {
		counts[input.Text()]++
		cnt++
	}
	if cnt > 1 {
		cntfiles[f.Name()] = cnt
	}
}
