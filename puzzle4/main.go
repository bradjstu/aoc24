package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	q3p1()
	q3p2()
}

func q3p2() {
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
	total := 0
	for i := range lines {
		for j := range lines[i] {
			next := scanSurroundingForMatchesP2(lines, i, j)
			total += next
		}
	}
	fmt.Println(total)

}

func scanSurroundingForMatchesP2(lines []string, i int, j int) int {
	ilen, jlen := len(lines), len(lines[0])
	imin, imax, jmin, jmax := i > 0, i < ilen-1, j > 0, j < jlen-1
	match1, match2 := "SAM", "MAS"
	count1, count2 := 0, 0
	if lines[i][j] != 'A' {
		return 0
	}
	rem := 0
	if imin && imax && jmin && jmax {
		rem = checkInDirection(match1, lines, i+1, j+1, -1, -1)
		count1 += rem
		count2 += checkInDirection(match2, lines, i+1, j+1, -1, -1)
		count1 += checkInDirection(match1, lines, i+1, j-1, -1, 1)
		count2 += checkInDirection(match2, lines, i+1, j-1, -1, 1)
	}
	if count1+count2 == 2 {
		return 1
	} else {
		return 0
	}
}

func q3p1() {
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
	total := 0
	for i := range lines {
		for j := range lines[i] {
			next := scanSurroundingForMatchesP1(lines, i, j)
			total += next
		}
	}
	fmt.Println(total)
}

func scanSurroundingForMatchesP1(lines []string, i int, j int) int {
	ilen, jlen := len(lines), len(lines[0])
	imin, imax, jmin, jmax := i > 2, i < ilen-3, j > 2, j < jlen-3
	match := "XMAS"
	total := 0
	if lines[i][j] != 'X' {
		return total
	}
	if imin {
		total += checkInDirection(match, lines, i, j, -1, 0)
	}
	if imin && jmin {
		total += checkInDirection(match, lines, i, j, -1, -1)
	}
	if imin && jmax {
		total += checkInDirection(match, lines, i, j, -1, 1)
	}
	if imax {
		total += checkInDirection(match, lines, i, j, 1, 0)
	}
	if imax && jmin {
		total += checkInDirection(match, lines, i, j, 1, -1)
	}
	if imax && jmax {
		total += checkInDirection(match, lines, i, j, 1, 1)
	}
	if jmin {
		total += checkInDirection(match, lines, i, j, 0, -1)
	}
	if jmax {
		total += checkInDirection(match, lines, i, j, 0, 1)
	}
	return total
}

func checkInDirection(match string, lines []string, i int, j int, idir int, jdir int) int {
	for index, letter := range []byte(match) {
		if !checkLetter(letter, lines, i, j, index*idir, index*jdir) {
			return 0
		}
	}
	return 1
}

func checkLetter(letter byte, lines []string, i int, j int, icount int, jcount int) bool {
	l2 := lines[i+icount][j+jcount]
	return l2 == letter
}
