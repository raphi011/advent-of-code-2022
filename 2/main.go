package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	rock    = 1
	paper   = 2
	scissor = 3

	loss = 'X'
	draw = 'Y'
	win  = 'Z'
)

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	acc := 0

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if len(line) != 4 {
			// 2 characters + whitespace + newline == 4
			panic("row must contain two chars")
		}

		p1 := line[0] - 'A' + 1

		if *exercise == 1 {
			p2 := line[2] - 'X' + 1

			acc += int(p2)

			if p1 == p2 {
				acc += 3
			} else if p2 == scissor && p1 == paper || p2 == paper && p1 == rock || p2 == rock && p1 == scissor {
				acc += 6
			}
		} else if *exercise == 2 {
			result := line[2]

			if result == draw {
				acc += 3
			} else if result == win {
				acc += 6
			}

			if p1 == rock && result == draw || p1 == paper && result == loss || p1 == scissor && result == win {
				acc += 1
			} else if p1 == rock && result == win || p1 == paper && result == draw || p1 == scissor && result == loss {
				acc += 2
			} else {
				acc += 3
			}
		} else {
			panic("unknown exercise")
		}
	}

	fmt.Printf("%d\n", acc)
}
