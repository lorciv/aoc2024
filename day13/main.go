package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

type machine struct {
	a, b  coord
	prize coord
}

func parse(r io.Reader) ([]machine, error) {
	var (
		cur      machine
		machines []machine
	)
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		line := scan.Text()

		if strings.HasPrefix(line, "Button A: ") {
			line = strings.TrimPrefix(line, "Button A: ")
			fmt.Sscanf(line, "X+%d, Y+%d", &cur.a.x, &cur.a.y)
		} else if strings.HasPrefix(line, "Button B: ") {
			line = strings.TrimPrefix(line, "Button B: ")
			fmt.Sscanf(line, "X+%d, Y+%d", &cur.b.x, &cur.b.y)
		} else if strings.HasPrefix(line, "Prize: ") {
			line = strings.TrimPrefix(line, "Prize: ")
			fmt.Sscanf(line, "X=%d, Y=%d", &cur.prize.x, &cur.prize.y)
		} else if line == "" {
			machines = append(machines, cur)
			cur = machine{}
		} else {
			return nil, errors.New("could not parse: unexpected line")
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}

	machines = append(machines, cur) // last one

	return machines, nil
}

func push(m machine, a, b int) (pos coord) {
	return coord{
		x: a*m.a.x + b*m.b.x,
		y: a*m.a.y + b*m.b.y,
	}
}

// tokens computes the minimum number of tokens to win a game on the machine.
// If win is not possible, it returns math.MaxInt.
func tokens(m machine) int {
	min := math.MaxInt
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			pos := push(m, a, b)
			if pos != m.prize {
				continue
			}
			c := 3*a + b
			if c < min {
				min = c
			}
		}
	}
	return min
}

func main() {
	machines, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, m := range machines {
		t := tokens(m)
		if t != math.MaxInt {
			total += tokens(m)
		}
	}
	fmt.Println(total)
}
