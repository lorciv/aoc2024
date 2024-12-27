package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}

func Safe(r []int) bool {
	// The levels are either all increasing or all decreasing.
	d0 := sign(r[0] - r[1])
	for i := 0; i < len(r)-1; i++ {
		d := sign(r[i] - r[i+1])
		if d != d0 {
			return false
		}
	}

	// Any two adjacent levels differ by at least one and at most three.
	for i := 0; i < len(r)-1; i++ {
		d := abs(r[i] - r[i+1])
		if d < 1 || d > 3 {
			return false
		}
	}

	return true
}

func main() {
	var records [][]int

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Fields(line)

		var record []int
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			record = append(record, n)
		}
		records = append(records, record)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, r := range records {
		if Safe(r) {
			count++
		}
	}
	fmt.Println(count)
}
