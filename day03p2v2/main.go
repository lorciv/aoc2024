package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// https://www.reddit.com/r/adventofcode/comments/1h5frsp/comment/m2wdvoo/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

var expression = regexp.MustCompile(`do\(\)|don't\(\)|mul\([\d]{1,3},[\d]{1,3}\)`)

func main() {
	var memory string
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		memory += line
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	do := true
	sum := 0

	instructions := expression.FindAllString(memory, -1)
	for _, t := range instructions {

		if t == "do()" {
			do = true
			continue
		}

		if t == "don't()" {
			do = false
			continue
		}

		if do {
			var a, b int
			fmt.Sscanf(t, "mul(%d,%d)", &a, &b)
			sum += a * b
		}
	}

	fmt.Println(sum)
}
