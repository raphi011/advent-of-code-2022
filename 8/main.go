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
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	var trees [][]int

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		treeline := make([]int, len(line))

		for i, h := range line {
			treeline[i], err = strconv.Atoi(string(h))
			if err != nil {
				panic(fmt.Sprintf("%c is not an integer: %v", h, err))
			}
		}
		trees = append(trees, treeline)
	}

	result := len(trees)*2 + len(trees[0])*2 - 4

	if *exercise == 2 {
		result = 1
	}

	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[0])-1; x++ {
			curHeight := trees[y][x]
			scenicScore := 1

			for _, dir := range []rune{'n', 'e', 's', 'w'} {
				var cx, cy int

				if *exercise == 1 {
					visible := true

					switch dir {
					case 'n':
						cx = x
						cy = 0
					case 'e':
						cx = len(trees[0]) - 1
						cy = y
					case 's':
						cx = x
						cy = len(trees) - 1
					case 'w':
						cy = y
						cx = 0
					}

					for cx != x || cy != y {
						if trees[cy][cx] >= curHeight {
							visible = false
							break
						}

						switch dir {
						case 'n':
							cy++
						case 'e':
							cx--
						case 's':
							cy--
						case 'w':
							cx++
						}
					}

					if visible {
						result++
						break
					}
				} else {
					next := func() {
						switch dir {
						case 'n':
							cy--
						case 'e':
							cx++
						case 's':
							cy++
						case 'w':
							cx--
						}
					}
					cx = x
					cy = y

					next()

					distance := 0

					for cx >= 0 && cx < len(trees[0]) && cy >= 0 && cy < len(trees) {
						distance++

						if trees[cy][cx] >= curHeight {
							break
						}

						next()
					}

					scenicScore *= distance
				}
			}

			if *exercise == 2 && scenicScore > result {
				result = scenicScore
			}
		}
	}

	fmt.Printf("%d\n", result)
}
