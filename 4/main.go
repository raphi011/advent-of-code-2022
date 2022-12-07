package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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

		var l1, u1, l2, u2 int
		if _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &l1, &u1, &l2, &u2); err != nil {
			panic(err)
		}

		if *exercise == 1 {
			if l1 <= l2 && u1 >= u2 || l2 <= l1 && u2 >= u1 {
				acc++
			}
		} else {
			if l2 < l1 {
				l1, u1, l2, u2 = l2, u2, l1, u1
			}

			if u1 >= l2 {
				acc++
			}
		}
	}

	fmt.Printf("%d\n", acc)
}
