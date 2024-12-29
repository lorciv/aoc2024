package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var expression = regexp.MustCompile(`^mul\([\d]{1,3},[\d]{1,3}\)`)

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

	sum := 0
	do := true

	for len(memory) > 0 {

		if strings.HasPrefix(memory, "do()") {
			do = true
			memory = strings.TrimPrefix(memory, "do()")
			continue
		}

		if strings.HasPrefix(memory, "don't()") {
			do = false
			memory = strings.TrimPrefix(memory, "don't()")
			continue
		}

		e := expression.FindString(memory)
		if e == "" {
			memory = memory[1:] // shift
			continue
		}
		if do {
			var a, b int
			fmt.Sscanf(e, "mul(%d,%d)", &a, &b)
			sum += a * b
		}
		memory = strings.TrimPrefix(memory, e)
	}

	fmt.Println(do, sum)
}
