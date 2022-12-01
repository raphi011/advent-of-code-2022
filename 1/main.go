package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	var maxc [3]int
	var acc int

	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		if line == "" {
			if acc > maxc[2] {
				maxc[0] = maxc[1]
				maxc[1] = maxc[2]
				maxc[2] = acc
			} else if acc > maxc[1] {
				maxc[0] = maxc[1]
				maxc[1] = acc
			} else if acc > maxc[0] {
				maxc[0] = acc
			}

			if err == io.EOF {
				break
			}

			acc = 0

			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("invalid number %s\n", line)
			os.Exit(1)
		}

		acc += calories
	}

	if *exercise == 1 {
		fmt.Printf("%d\n", maxc[2])
	} else {
		fmt.Printf("%v = %d\n", maxc, maxc[0]+maxc[1]+maxc[2])
	}
}
