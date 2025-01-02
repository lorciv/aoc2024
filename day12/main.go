package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	garden     [][]string
	rows, cols int
)

func read() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		row := strings.Split(line, "")
		garden = append(garden, row)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	rows = len(garden)
	cols = len(garden[0])
}

func show() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(garden[i][j] + " ")
		}
		fmt.Println()
	}
}

type coord struct {
	x, y int
}

func neighbors(c coord) []coord {
	dirs := []coord{
		{-1, 0}, // up
		{1, 0},  // down
		{0, 1},  // right
		{0, -1}, // left
	}

	var result []coord
	for _, dir := range dirs {
		d := coord{c.x + dir.x, c.y + dir.y}

		// skip if out of grid
		if d.x < 0 || d.x >= len(garden) || d.y < 0 || d.y >= len(garden[0]) {
			continue
		}

		// skip if different plant
		if garden[d.x][d.y] != garden[c.x][c.y] {
			continue
		}

		result = append(result, d)
	}
	return result
}

func region(start coord) []coord {
	queue := []coord{start}
	visited := map[coord]bool{start: true}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, n := range neighbors(cur) {
			if !visited[n] {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}

	var result []coord
	for c := range visited {
		result = append(result, c)
	}
	return result
}

func explore() [][]coord {
	var regions [][]coord

	visited := make(map[coord]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			c := coord{i, j}
			if visited[c] {
				continue
			}
			reg := region(c)
			regions = append(regions, reg)
			for _, c := range reg {
				visited[c] = true
			}
		}
	}

	return regions
}

func area(region []coord) int {
	return len(region)
}

func perimeter(region []coord) int {
	// initial guess of 4 fences per square
	// then remove fences between adjacent squares
	total := 0
	for _, c := range region {
		total += 4 - len(neighbors(c))
	}
	return total
}

func price(region []coord) int {
	return area(region) * perimeter(region)
}

func main() {
	read()
	// show()

	regions := explore()

	total := 0
	for _, r := range regions {
		total += price(r)
	}
	fmt.Println(total)
}
