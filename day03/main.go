package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var expression = regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)

func main() {
	sum := 0

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()

		matches := expression.FindAllString(line, -1)
		for _, m := range matches {
			var a, b int
			fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
			sum += a * b
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
