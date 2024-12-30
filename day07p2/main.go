package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	test    int
	numbers []int
}

func evaluate(numbers []int, operators string) (int, error) {
	if len(operators) != len(numbers)-1 {
		return 0, errors.New("incompatible lengths of numbers and operators")
	}

	result := 0
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			result = numbers[i]
			continue
		}

		op := operators[i-1]
		switch op {
		case '+':
			result += numbers[i]
		case '*':
			result *= numbers[i]
		case '|':
			s := strconv.Itoa(result) + strconv.Itoa(numbers[i])
			result, _ = strconv.Atoi(s)
		default:
			return 0, fmt.Errorf("unsupported operator: %c", op)
		}
	}
	return result, nil
}

func perm(dict string, n int) []string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return strings.Split(dict, "")
	}

	heads := perm(dict, 1)
	tails := perm(dict, n-1)

	var result []string
	for _, h := range heads {
		for _, t := range tails {
			result = append(result, h+t)
		}
	}
	return result
}

func parse(line string) equation {
	split := strings.Split(line, ":")
	test, _ := strconv.Atoi(split[0])
	var numbers []int
	for _, f := range strings.Fields(split[1]) {
		n, _ := strconv.Atoi(f)
		numbers = append(numbers, n)
	}
	return equation{test: test, numbers: numbers}
}

func main() {
	var equations []equation
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		e := parse(line)
		equations = append(equations, e)
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, e := range equations {
		operators := perm("+*|", len(e.numbers)-1)
		for _, p := range operators {
			got, err := evaluate(e.numbers, p)
			if err != nil {
				log.Fatal(err)
			}
			if got == e.test {
				total += e.test
				break
			}
		}
	}
	fmt.Println(total)
}
