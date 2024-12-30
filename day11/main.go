package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func read() ([]int, error) {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	line := string(raw)
	fields := strings.Fields(line)
	stones := make([]int, len(fields))
	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		stones[i] = n
	}
	return stones, nil
}

func blink(cur []int) []int {
	var next []int
	for _, c := range cur {
		if c == 0 {
			next = append(next, 1)
		} else if s := strconv.Itoa(c); len(s)%2 == 0 {
			mid := len(s) / 2
			s1, s2 := s[:mid], s[mid:]
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			next = append(next, n1, n2)
		} else {
			next = append(next, c*2024)
		}
	}
	return next
}

func main() {
	n := flag.Int("n", 25, "number of blinks")
	flag.Parse()

	cur, err := read()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(cur)
	for i := 0; i < *n; i++ {
		cur = blink(cur)
		// fmt.Println(cur)
	}
	fmt.Println(len(cur))
}
