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

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func main() {
	var list1, list2 []int

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Fields(line)

		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, n1)

		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		list2 = append(list2, n2)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i := 0; i < len(list1); i++ {
		d := list1[i] - list2[i]
		sum += abs(d)
	}
	fmt.Println(sum)
}
