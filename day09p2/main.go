package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const free = -1

func parse(line string) []int {
	var memory []int

	id := 0
	for i, s := range strings.Split(line, "") {
		n, _ := strconv.Atoi(s)

		block := free
		if i%2 == 0 {
			block = id
			id++
		}

		for i := 0; i < n; i++ {
			memory = append(memory, block)
		}
	}

	return memory
}

func show(memory []int) {
	for _, block := range memory {
		if block == free {
			fmt.Print(".")
			continue
		}
		fmt.Print(block)
	}
	fmt.Println()
}

// defrag moves one used block to the left-most free space, if possible.
// It returns true if at least one block was moved.
func defrag(memory []int) bool {
	var src int
	for src = len(memory) - 1; src >= 0; src-- {
		if memory[src] != free {
			break
		}
	}

	var dest int
	for dest = 0; dest < len(memory); dest++ {
		if memory[dest] == free {
			break
		}
	}

	if dest < src {
		memory[dest] = memory[src]
		memory[src] = free
		return true
	}

	return false
}

func checksum(memory []int) int {
	result := 0
	for i, block := range memory {
		if block == free {
			continue
		}
		result += i * block
	}
	return result
}

type segment struct {
	start, end int
	id         int
}

func length(s segment) int {
	return s.end - s.start
}

var seen = make(map[int]bool)

func move(memory []int) bool {
	var segments []segment
	i := 0
	for i < len(memory) {
		j := i + 1
		for j < len(memory) && memory[j] == memory[i] {
			j++
		}
		segments = append(segments, segment{start: i, end: j, id: memory[i]})
		i = j
	}

	for i := len(segments) - 1; i >= 0; i-- {
		if segments[i].id == free {
			continue
		}
		file := segments[i]

		if seen[file.id] {
			continue
		}
		seen[file.id] = true

		for j := 0; j < i; j++ {
			if segments[j].id != free {
				continue
			}
			space := segments[j]

			if length(space) < length(file) {
				continue
			}

			// move file to free space
			src := memory[file.start:file.end]
			copy(memory[space.start:space.end], src)
			for k := range src {
				src[k] = -1
			}
			return true
		}
	}
	return false
}

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	memory := parse(string(raw))

	// show(memory)

	for move(memory) {
		// show(memory)
	}

	fmt.Println(checksum(memory))

}
