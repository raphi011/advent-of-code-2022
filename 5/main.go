package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type stack []rune

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	var stacks []stack

	instructions := false

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else if line == "\n" {
			for i := range stacks {
				// needed since we're parsing the stacks in the wrong order (from the top)
				stacks[i].reverse()
			}

			instructions = true

			continue
		}

		runes := []rune(line)

		if !instructions {
			// parse the stacks
			for i := 1; i < len(runes); i += 4 {

				if line[i] == ' ' {
					continue
				}

				cur := i / 4

				for len(stacks) < cur+1 {
					stacks = append(stacks, stack{})
				}

				stacks[cur] = append(stacks[cur], runes[i])
			}
		} else {
			// run the moves
			var count, from, to int
			_, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
			if err != nil {
				panic("invalid move")
			}

			if *exercise == 1 {
				for i := 0; i < count; i++ {
					if stacks[from-1].isEmpty() {
						continue
					}

					stacks[to-1].push(stacks[from-1].pop())
				}
			} else {
				stacks[to-1].add(stacks[from-1].remove(count))
			}
		}
	}

	var top []rune

	for _, r := range stacks {
		top = append(top, r.peek())
	}

	fmt.Printf("%s\n", string(top))
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) reverse() {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i], (*s)[len(*s)-i-1] = (*s)[len(*s)-i-1], (*s)[i]
	}
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) peek() rune {
	if s.isEmpty() {
		return rune(0)
	}

	return (*s)[len(*s)-1]
}

func (s *stack) remove(c int) stack {
	size := len(*s)
	if c > size {
		c = size
	}

	elems := (*s)[size-c:]

	*s = (*s)[:size-c]

	return elems
}

func (s *stack) add(n stack) {
	for _, e := range n {
		(*s) = append(*s, e)
	}
}

func (s *stack) pop() rune {
	if s.isEmpty() {
		return rune(0)
	}
	i := len(*s) - 1
	r := (*s)[i]
	*s = (*s)[:i]

	return r
}
