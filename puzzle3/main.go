package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// q3p1()
	q3p2()
}

func q3p2() {
	do, dont := "do()", "don't()"
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(in[:])
	matches := regex.FindAllStringSubmatch(str, -1)

	active := true
	total := 0
	for _, match := range matches {
		if do == match[0] {
			active = true
		} else if dont == match[0] {
			active = false
		} else if active {
			i1, _ := strconv.Atoi(match[1])
			i2, _ := strconv.Atoi(match[2])
			multiplied := i1 * i2
			total += multiplied
		}
	}
	fmt.Println(total)
}

func q3p1() {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(in[:])

	total := 0
	matches := mulRegex.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		// fmt.Println(match[1] + " " + match[2])
		i1, _ := strconv.Atoi(match[1])
		i2, _ := strconv.Atoi(match[2])
		multiplied := i1 * i2
		total += multiplied
	}
	fmt.Println(total)
}
