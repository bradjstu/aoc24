package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	q7p1()
	q7p2()
}

func q7p2() {
	lines := readInLines()

	total := int64(0)
	for _, line := range lines {
		result, in := readLine(line)
		if rec2(in, result, 1, in[0]) {
			total += result
		}
	}
	fmt.Println(total)
}

func q7p1() {
	lines := readInLines()

	total := int64(0)
	for _, line := range lines {
		result, in := readLine(line)
		if rec(in, result, 1, in[0]) {
			total += result
		}
	}
	fmt.Println(total)
}

func rec2(array []int64, result int64, index int, cumulative int64) bool {
	if index == len(array) {
		return result == cumulative
	}

	number := array[index]
	result1 := rec2(array, result, index+1, cumulative+number)
	result2 := rec2(array, result, index+1, cumulative*number)
	result3 := rec2(array, result, index+1, conc(cumulative, number))

	return result1 || result2 || result3
}

func conc(pre int64, post int64) int64 {
	num := strconv.FormatInt(pre, 10) + strconv.FormatInt(post, 10)
	i64, _ := strconv.ParseInt(num, 10, 64)
	return i64
}

func rec(array []int64, result int64, index int, cumulative int64) bool {
	if index == len(array) {
		return result == cumulative
	}

	number := array[index]
	result1 := rec(array, result, index+1, cumulative+number)
	result2 := rec(array, result, index+1, cumulative*number)

	return result1 || result2
}

func readInLines() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func readLine(line string) (int64, []int64) {
	split := strings.Split(line, ":")
	i64, _ := strconv.ParseInt(split[0], 10, 64)

	strs := strings.Fields(split[1])
	values := []int64{}
	for _, str := range strs {
		v, _ := strconv.ParseInt(str, 10, 64)
		values = append(values, v)
	}
	return i64, values
}
