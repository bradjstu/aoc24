package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// print(q1p1())
	// print(q1p2())
	print(q1p2a2())
}

func q1p2a2() int {
	l1, l2 := toSorted()
	total, p1, p2, count := 0, 0, 0, 0

	for p1 < len(l1) && p2 < len(l2) {
		if l1[p1] > l2[p2] {
			count = 0
			p2++
			continue
		}
		if p2 < len(l2)-1 && l2[p2] == l2[p2+1] {
			count++
			p2++
			continue
		}
		if l1[p1] < l2[p2] {
			p1++
			continue
		}
		if l1[p1] == l2[p2] {
			total += l1[p1] * (count + 1)
			p1++
			continue
		}
	}
	return total
}

func q1p2() int {
	l1, l2 := toUnSorted()
	total := 0
	for _, element := range l1 {
		count := 0
		for _, element2 := range l2 {
			if element == element2 {
				count++
			}
		}
		total += count * element
	}
	return total
}

func q1p1() int {
	l1, l2 := toSorted()
	total := 0
	for index := range len(l1) {
		diff := absDiffInt(l1[index], l2[index])
		total += diff
	}
	return total
}

func toSorted() ([]int, []int) {
	l1, l2 := toUnSorted()
	sort.Ints(l1)
	sort.Ints(l2)
	return l1, l2
}

func toUnSorted() ([]int, []int) {
	in, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(in[:])

	var l1, l2 []int
	split := strings.Fields(str)

	for index, element := range split {
		if index%2 == 0 {
			s, _ := strconv.Atoi(element)
			l1 = append(l1, s)
		} else {
			s, _ := strconv.Atoi(element)
			l2 = append(l2, s)
		}
	}
	return l1, l2
}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
