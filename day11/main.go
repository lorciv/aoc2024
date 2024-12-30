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

// func blink(cur []int) []int {
// 	var next []int
// 	for _, c := range cur {
// 		if c == 0 {
// 			next = append(next, 1)
// 		} else if s := strconv.Itoa(c); len(s)%2 == 0 {
// 			mid := len(s) / 2
// 			s1, s2 := s[:mid], s[mid:]
// 			n1, _ := strconv.Atoi(s1)
// 			n2, _ := strconv.Atoi(s2)
// 			next = append(next, n1, n2)
// 		} else {
// 			next = append(next, c*2024)
// 		}
// 	}
// 	return next
// }

type input struct {
	stone, blinks int
}

func count(stone, blinks int, memory map[input]int) int {
	if blinks == 0 {
		return 1
	}

	if result, ok := memory[input{stone, blinks}]; ok {
		return result
	}

	if stone == 0 {
		result := count(1, blinks-1, memory)
		memory[input{1, blinks - 1}] = result
		return result
	}
	if s := strconv.Itoa(stone); len(s)%2 == 0 {
		mid := len(s) / 2
		s1, s2 := s[:mid], s[mid:]
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)

		res1 := count(n1, blinks-1, memory)
		memory[input{n1, blinks - 1}] = res1

		res2 := count(n2, blinks-1, memory)
		memory[input{n2, blinks - 1}] = res2

		return res1 + res2
	}

	res := count(stone*2024, blinks-1, memory)
	memory[input{stone * 2024, blinks - 1}] = res
	return res
}

func main() {
	n := flag.Int("n", 25, "number of blinks")
	flag.Parse()

	cur, err := read()
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, stone := range cur {
		sum += count(stone, *n, make(map[input]int))
	}
	fmt.Println(sum)
}
