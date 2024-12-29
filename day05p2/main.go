package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	before, after int
}

type update []int

// valid checks an update u against the rules.
func valid(u update, rules []rule) bool {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			for _, r := range rules {
				if u[i] == r.after && u[j] == r.before {
					return false
				}
			}
		}
	}
	return true
}

func fix(u update, rules []rule) update {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			for _, r := range rules {
				if u[i] == r.after && u[j] == r.before {
					u[i], u[j] = u[j], u[i] // swap
				}
			}
		}
	}
	return u
}

func middle(u update) int {
	return u[len(u)/2]
}

func main() {
	var (
		rules   []rule
		updates []update
	)

	scan := bufio.NewScanner(os.Stdin)
	// section 1
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break
		}

		var r rule
		fmt.Sscanf(line, "%d|%d", &r.before, &r.after)
		rules = append(rules, r)
	}
	// section 2
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, ",")

		var u update
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			u = append(u, n)
		}
		updates = append(updates, u)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, u := range updates {
		if !valid(u, rules) {
			v := fix(u, rules)
			sum += middle(v)
		}
	}
	fmt.Println(sum)
}
