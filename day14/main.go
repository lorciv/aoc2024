package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

type robot struct {
	pos, vel coord
}

func parse() ([]robot, error) {
	var robots []robot

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		var r robot
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.pos.x, &r.pos.y, &r.vel.x, &r.vel.y)
		if err != nil {
			return nil, err
		}
		robots = append(robots, r)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}

func show(robots []robot) {
	count := make(map[coord]int)
	for _, r := range robots {
		count[r.pos]++
	}
	for y := 0; y < *height; y++ {
		for x := 0; x < *width; x++ {
			c := count[coord{x, y}]
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", c)
			}
		}
		fmt.Println()
	}
}

func move(r robot) robot {
	s := robot{
		pos: coord{
			x: ((r.pos.x + r.vel.x) + *width) % *width,
			y: ((r.pos.y + r.vel.y) + *height) % *height,
		},
		vel: r.vel,
	}
	return s
}

func safety(robots []robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		if r.pos.x < *width/2 && r.pos.y < *height/2 {
			q1++
		}
		if r.pos.x > *width/2 && r.pos.y < *height/2 {
			q2++
		}
		if r.pos.x > *width/2 && r.pos.y > *height/2 {
			q3++
		}
		if r.pos.x < *width/2 && r.pos.y > *height/2 {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

var (
	width  = flag.Int("w", 11, "width of space")
	height = flag.Int("h", 7, "height of space")
)

func main() {
	flag.Parse()

	robots, err := parse()
	if err != nil {
		log.Fatal(err)
	}

	for range 100 {
		for i := range robots {
			robots[i] = move(robots[i])
		}
	}

	fmt.Println(safety(robots))
}
