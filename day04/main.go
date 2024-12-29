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

var directions = []coord{
	{-1, 0},  // up
	{1, 0},   // down
	{0, 1},   // right
	{0, -1},  // left
	{-1, 1},  // up-right
	{1, 1},   // down-right
	{-1, -1}, // up-left
	{1, -1},  // down-left
}

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

			for _, d := range directions {
				c1 := cell(grid, cur)
				c2 := cell(grid, next(cur, d))
				c3 := cell(grid, next(next(cur, d), d))
				c4 := cell(grid, next(next(next(cur, d), d), d))
				word := c1 + c2 + c3 + c4

				if word == "XMAS" {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
