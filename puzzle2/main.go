package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// q2p1()
	q2p2()
}

func q2p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isSafe(line) {
			count++
		} else {
			split := strings.Fields(line)
			for idx := range split {
				tmp := make([]string, 0, len(split)-1)
				for i := range split {
					if i != idx {
						tmp = append(tmp, split[i])
					}
				}
				if isSafe(strings.Join(tmp, " ")) {
					count++
					break
				}
			}
		}
	}
	print(count)
}

func q2p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isSafe(scanner.Text()) {
			count++
		}
	}
	print(count)
}

func isSafe(line string) bool {
	split := strings.Fields(line)
	first, _ := strconv.Atoi(split[0])
	prev, _ := strconv.Atoi(split[1])
	diff := first - prev
	absDiff := absDiffInt(first, prev)
	isPos := diff > 0
	if absDiff < 1 || absDiff > 3 {
		return false
	}
	for i := 2; i < len(split); i++ {
		next, _ := strconv.Atoi(split[i])
		diff = prev - next
		absDiff = absDiffInt(prev, next)
		if absDiff < 1 || absDiff > 3 || ((diff > 0) != isPos) {
			return false
		}
		prev = next
	}
	return true

}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
