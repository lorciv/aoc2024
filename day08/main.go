package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// consider all antennas with the same frequency
// consider all possible pairs
// mark the antinodes on the map
// finally, count the antinodes

type coord struct {
	row, col int
}

type antenna struct {
	freq string
	pos  coord
}

func antinodes(a, b coord) []coord {
	drow, dcol := b.row-a.row, b.col-a.col
	n1 := coord{
		row: a.row - (drow),
		col: a.col - (dcol),
	}
	n2 := coord{
		row: b.row + (drow),
		col: b.col + (dcol),
	}
	return []coord{n1, n2}
}

var rows, cols int

func valid(c coord) bool {
	return c.row >= 0 && c.row < rows && c.col >= 0 && c.col < cols
}

func main() {
	antennas := make(map[string][]antenna)

	i := 0
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, "")
		for j, s := range split {
			if s != "." {
				a := antenna{freq: s, pos: coord{i, j}}
				antennas[a.freq] = append(antennas[a.freq], a)
			}
		}
		i++

		cols = len(split)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	rows = i

	set := make(map[coord]bool)
	for _, subset := range antennas {
		for i, a := range subset {
			for _, b := range subset[i+1:] {
				for _, n := range antinodes(a.pos, b.pos) {
					if valid(n) {
						set[n] = true
					}
				}
			}
		}
	}
	for c := range set {
		fmt.Println(c)
	}
	fmt.Println(len(set))
}
