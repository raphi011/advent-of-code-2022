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

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	if *exercise == 1 {
		first()
	} else {
		second()
	}
}

func first() {
	reader := bufio.NewReader(os.Stdin)

	acc := 0

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		var left, right uint64

		for i, c := range strings.TrimSpace(line) {
			var exp uint64
			if c >= 'a' && c <= 'z' {
				exp = uint64(c) - uint64('a')
			} else {
				exp = uint64(c) - uint64('A') + 26
			}

			if i < len(line)/2 {
				left |= uint64(math.Pow(2, float64(exp)))
			} else {
				right |= uint64(math.Pow(2, float64(exp)))
			}

		}

		overlap := left & right

		val := int(math.Log2(float64(overlap))) + 1

		acc += val
	}

	fmt.Printf("%d\n", acc)
}

func second() {
	reader := bufio.NewReader(os.Stdin)

	acc := 0

	for {
		var overlap uint64

		for i := 0; i < 3; i++ {
			line, err := reader.ReadString('\n')

			if err == io.EOF && i == 0 {
				fmt.Printf("%d\n", acc)
				return
			} else if err != nil {
				panic(err)
			}

			var charMap uint64

			for _, c := range strings.TrimSpace(line) {
				var exp uint64

				if c >= 'a' && c <= 'z' {
					exp = uint64(c) - uint64('a')
				} else {
					exp = uint64(c) - uint64('A') + 26
				}

				charMap |= uint64(math.Pow(2, float64(exp)))
			}

			if i == 0 {
				overlap |= charMap
			} else {
				overlap &= charMap
			}
		}

		val := int(math.Log2(float64(overlap))) + 1

		acc += val
	}
}
