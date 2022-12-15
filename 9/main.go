package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	visitMap := map[string]bool{"(0,0)": true}

	length := 2

	if *exercise == 2 {
		length = 10
	}

	coords := make([]coords, length)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		result := strings.Split(line, " ")
		dir := result[0]
		steps, err := strconv.Atoi(result[1])
		if err != nil {
			panic("steps is not a number")
		}

		for i := 0; i < steps; i++ {
			switch dir {
			case "U":
				coords[0].y--
			case "R":
				coords[0].x++
			case "L":
				coords[0].x--
			case "D":
				coords[0].y++
			}

			for i := 1; i < len(coords); i++ {
				prev := coords[i-1]
				c := &coords[i]

				dist := math.Abs(float64(prev.x-c.x)) + math.Abs(float64(prev.y-c.y))

				if (prev.x != c.x && prev.y != c.y && dist < 3) || dist < 2 {
					continue
				}

				if prev.x == c.x || prev.y != c.y {
					diff := int(math.Abs(float64(prev.y)-float64(c.y))) / (prev.y - c.y)
					c.y += diff
				}

				if prev.y == c.y || prev.x != c.x {
					diff := int(math.Abs(float64(prev.x)-float64(c.x))) / (prev.x - c.x)
					c.x += diff
				}
			}

			tail := coords[len(coords)-1]

			visitMap[fmt.Sprintf("(%d,%d)", tail.x, tail.y)] = true
		}

	}

	fmt.Printf("%d\n", len(visitMap))
}
