package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// horizontal, vertical, diagonal, written backwards, or even overlapping other words

type coord struct {
	row, col int
}

var (
	dirUp        = coord{-1, 0}
	dirDown      = coord{1, 0}
	dirRight     = coord{0, 1}
	dirLeft      = coord{0, -1}
	dirUpRight   = coord{-1, 1}
	dirDownRight = coord{1, 1}
	dirUpLeft    = coord{-1, -1}
	dirDownLeft  = coord{1, -1}
)

var directions = []coord{dirUp, dirDown, dirRight, dirLeft, dirUpRight, dirDownRight, dirUpLeft, dirDownLeft}

func next(c, dir coord) coord {
	return coord{
		row: c.row + dir.row,
		col: c.col + dir.col,
	}
}

func cell(grid [][]string, c coord) string {
	if c.row < 0 || c.row >= len(grid) || c.col < 0 || c.col > len(grid[c.row])-1 {
		return "."
	}
	return grid[c.row][c.col]
}

func word(grid [][]string, start, dir coord) string {
	w := ""
	cur := start
	for i := 0; i < 3; i++ {
		w += cell(grid, cur)
		cur = next(cur, dir)
	}
	return w
}

func main() {
	var grid [][]string

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0

	for i := range grid {
		for j := range grid[i] {
			cur := coord{i, j}

			w1 := word(grid, next(cur, dirUpLeft), dirDownRight)
			if w1 != "MAS" && w1 != "SAM" {
				continue
			}

			w2 := word(grid, next(cur, dirUpRight), dirDownLeft)
			if w2 != "MAS" && w2 != "SAM" {
				continue
			}

			count++
		}
	}

	fmt.Println(count)
}
