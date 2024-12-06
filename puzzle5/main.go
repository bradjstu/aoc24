package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	q5p1()
	q5p2()
}

func q5p2() {
	regex := regexp.MustCompile(`(\d+)\|(\d+)`)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var linesRules []string
	var linesUpdates []string
	scanner := bufio.NewScanner(file)
	updates := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			updates = true
			continue
		}
		if updates {
			linesUpdates = append(linesUpdates, scanner.Text())
		} else {
			linesRules = append(linesRules, scanner.Text())
		}
	}
	rules := make(map[int][]int)

	for _, line := range linesRules {
		match := regex.FindStringSubmatch(line)
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		values := rules[n1]
		if values == nil {
			values = []int{n2}
			rules[n1] = values
		} else {
			values = append(values, n2)
			rules[n1] = values
		}
	}
	total, count := 0, 0
	for _, line := range linesUpdates {
		split := strings.Split(line, ",")
		v := []int{}
		for _, field := range split {
			n, _ := strconv.Atoi(field)
			v = append(v, n)
		}
		if !isOrdered(rules, v) {
			v = order(rules, v)
			index := (len(v) - 1) / 2
			total += v[index]
			count++
		}
	}
	fmt.Println(total)
	fmt.Println(count)
}

func order(rules map[int][]int, values []int) []int {
	tmp := values
	// Assuming none impossible to order, if impossible then cache already seen combinations, to break out
	for !isOrdered(rules, tmp) {
		for _, value := range values {
			r := rules[value]
			for _, swap := range r {
				swapIndex := -1
				for i := 0; i < len(tmp); i++ {
					if tmp[i] == value {
						if swapIndex > -1 {
							tmp[i] = swap
							tmp[swapIndex] = value
						}
						break
					} else if tmp[i] == swap {
						swapIndex = i
					}
				}
			}
		}
	}
	if !isOrdered(rules, tmp) {
		fmt.Println(tmp)
	}
	return tmp
}

func q5p1() {
	regex := regexp.MustCompile(`(\d+)\|(\d+)`)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var linesRules []string
	var linesUpdates []string
	scanner := bufio.NewScanner(file)
	updates := false
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			updates = true
			continue
		}
		if updates {
			linesUpdates = append(linesUpdates, scanner.Text())
		} else {
			linesRules = append(linesRules, scanner.Text())
		}
	}
	rules := make(map[int][]int)

	for _, line := range linesRules {
		match := regex.FindStringSubmatch(line)
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		values := rules[n1]
		if values == nil {
			values = []int{n2}
			rules[n1] = values
		} else {
			values = append(values, n2)
			rules[n1] = values
		}
	}
	total, count := 0, 0
	for _, line := range linesUpdates {
		split := strings.Split(line, ",")
		v := []int{}
		for _, field := range split {
			n, _ := strconv.Atoi(field)
			v = append(v, n)
		}
		if isOrdered(rules, v) {
			index := (len(v) - 1) / 2
			total += v[index]
			count++
		}

	}
	fmt.Println(total)
	fmt.Println(count)
}

func isOrdered(rules map[int][]int, values []int) bool {
	for idx, value := range values {
		r := rules[value]
		for i := idx; i > 0; i-- {
			if slices.Contains(r, values[i]) {
				return false
			}
		}
	}
	return true
}
