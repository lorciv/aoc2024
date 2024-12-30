package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// a hiking TRAIL is any path that
// - starts at height 0
// - ends at height 9
// - always increases by a height of exactly 1 at each step
// - never include diagonal steps - only up, down, left, or right

// a TRAILHEAD is any position that starts one or more hiking trails

// a trailhead's SCORE is the number of 9-height positions reachable from that trailhead via a hiking trail

var (
	grid       [][]int
	rows, cols int
)

func read() error {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, "")
		row := make([]int, len(split))
		for i, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			row[i] = n
		}
		grid = append(grid, row)
		cols = len(row)
	}
	if err := scan.Err(); err != nil {
		return err
	}
	rows = len(grid)
	return nil
}

// func show() {
// 	for i := range grid {
// 		for j := range grid[i] {
// 			fmt.Printf("%d ", grid[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

type coord struct {
	row, col int
}

func trailheads() []coord {
	var result []coord
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				result = append(result, coord{i, j})
			}
		}
	}
	return result
}

func neighbors(c coord) []coord {
	dirs := []coord{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var result []coord
	for _, d := range dirs {
		r := coord{
			row: c.row + d.row,
			col: c.col + d.col,
		}
		// avoid cells out of grid
		if r.row < 0 || r.row >= rows || r.col < 0 || r.col >= cols {
			continue
		}
		// avoid cells going downhill
		if grid[r.row][r.col] != grid[c.row][c.col]+1 {
			continue
		}
		result = append(result, r)
	}
	return result
}

func score(trailhead coord) int {
	count := 0

	queue := []coord{trailhead}
	visited := map[coord]bool{trailhead: true}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if grid[cur.row][cur.col] == 9 {
			count++
		}

		for _, n := range neighbors(cur) {
			if !visited[n] {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}

	return count
}

type trail struct {
	head coord
	tail *trail
}

func visited(t *trail, c coord) bool {
	if t == nil {
		return false
	}
	if t.head == c {
		return true
	}
	return visited(t.tail, c)
}

func rating(trailhead coord) int {
	count := 0

	queue := []*trail{{head: trailhead}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if grid[cur.head.row][cur.head.col] == 9 {
			count++
			continue
		}

		for _, n := range neighbors(cur.head) {
			if !visited(cur, n) {
				queue = append(queue, &trail{head: n, tail: cur})
			}
		}
	}

	return count
}

func main() {
	if err := read(); err != nil {
		log.Fatal(err)
	}

	totalScore := 0
	for _, h := range trailheads() {
		totalScore += score(h)
	}
	fmt.Println(totalScore)

	totalRating := 0
	for _, h := range trailheads() {
		totalRating += rating(h)
	}
	fmt.Println(totalRating)
}
