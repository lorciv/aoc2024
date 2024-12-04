package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var l1, l2 []int

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Fields(line)

		n, _ := strconv.Atoi(fields[0])
		l1 = append(l1, n)

		m, _ := strconv.Atoi(fields[1])
		l2 = append(l2, m)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(l1)
	sort.Ints(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		d := l1[i] - l2[i]
		sum += abs(d)
	}
	fmt.Println(sum)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
