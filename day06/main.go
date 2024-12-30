package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	grid       [][]string
	rows, cols int
	over       bool = false
)

// func show() {
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			fmt.Print(grid[i][j] + " ")
// 		}
// 		fmt.Println()
// 	}
// 	for j := 0; j < cols; j++ {
// 		fmt.Print("- ")
// 	}
// 	fmt.Println()
// }

// cell returns the content of the cell at row i, col j.
// If the cell is outside the grid, it returns the empty string.
func cell(i, j int) string {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return ""
	}
	return grid[i][j]
}

func step() {

	if over {
		return
	}

	// turn 90 degrees
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch cell(i, j) {
			case "^":
				next := cell(i-1, j)
				if next == "#" {
					grid[i][j] = ">"
				}
			case ">":
				next := cell(i, j+1)
				if next == "#" {
					grid[i][j] = "v"
				}
			case "v":
				next := cell(i+1, j)
				if next == "#" {
					grid[i][j] = "<"
				}
			case "<":
				next := cell(i, j-1)
				if next == "#" {
					grid[i][j] = "^"
				}
			}
		}
	}

	// move
loop:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch cell(i, j) {
			case "^":
				grid[i][j] = "X"
				next := cell(i-1, j)
				if next == "" {
					over = true
					break loop
				}
				grid[i-1][j] = "^"
				break loop
			case ">":
				grid[i][j] = "X"
				next := cell(i, j+1)
				if next == "" {
					over = true
					break loop
				}
				grid[i][j+1] = ">"
				break loop
			case "v":
				grid[i][j] = "X"
				next := cell(i+1, j)
				if next == "" {
					over = true
					break loop
				}
				grid[i+1][j] = "v"
				break loop
			case "<":
				grid[i][j] = "X"
				next := cell(i, j-1)
				if next == "" {
					over = true
					break loop
				}
				grid[i][j-1] = "<"
				break loop
			}
		}
	}

}

func visited() int {
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "X" {
				count++
			}
		}
	}
	return count
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, "")

		grid = append(grid, split)
		rows++
		cols = len(grid)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	for !over {
		step()
	}

	fmt.Println(visited())
}
