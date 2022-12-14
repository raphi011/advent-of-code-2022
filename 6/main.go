package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	runes := []rune(line)

	distinctMarkers := 4
	if *exercise == 2 {
		distinctMarkers = 14
	}

	for i := distinctMarkers; i < len(runes); i++ {
		set := make(map[rune]bool)

		unique := true

		for _, r := range runes[i-distinctMarkers : i] {
			if set[r] {
				unique = false
				break
			}

			set[r] = true
		}

		if unique {
			fmt.Printf("%d\n", i)
			os.Exit(0)
		}
	}

	fmt.Println("-1")
}
