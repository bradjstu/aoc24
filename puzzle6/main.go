package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	q6p1()
	q6p2()
}

const block byte = '#'
const been byte = 'X'
const free byte = '.'
const up byte = '^'
const down byte = 'v'
const left byte = '<'
const right byte = '>'

func q6p2() {
	lines := readInGrid()

	imax, jmax := len(lines)-1, len(lines[0])-1
	guardi, guardj := locateGuard(lines)

	toBlock := make([][]bool, imax+1)
	beenTo := make([][]int, imax+1)
	tmpLines := make([][]byte, imax+1)
	for idx := range lines {
		toBlock[idx] = make([]bool, jmax+1)
		beenTo[idx] = make([]int, jmax+1)
		tmpLines[idx] = make([]byte, jmax+1)
	}

	i, j := guardi, guardj
	currentDir := lines[guardi][guardj]
	for ii := range lines {
		for jj := range lines[0] {
			tmpLines[ii][jj] = lines[ii][jj]
		}
	}
	for i > 0 && j > 0 && i < imax && j < jmax {
		ii, jj := direction(currentDir)
		if checkNextSquare(lines, i, j, ii, jj) {
			tmpLines[i][j] = been
			i += ii
			j += jj
			toBlock[i][j] = true
		} else {
			currentDir = rotate(currentDir)
		}
	}

	total := 0
	for bi := 0; bi < len(toBlock); bi++ {
		for bj := 0; bj < len(toBlock[0]); bj++ {
			if !toBlock[bi][bj] {
				continue
			}
			i, j = guardi, guardj
			currentDir = lines[guardi][guardj]
			for ii := range lines {
				for jj := range lines[0] {
					beenTo[ii][jj] = 0
					tmpLines[ii][jj] = lines[ii][jj]
				}
			}
			tmpLines[bi][bj] = block
			for i > 0 && j > 0 && i < imax && j < jmax {
				ii, jj := direction(currentDir)
				if checkNextSquare(tmpLines, i, j, ii, jj) {
					tmpLines[i][j] = been
					beenTo[i][j]++
					if beenTo[i][j] > 4 {
						total++
						break
					}
					i += ii
					j += jj
				} else {
					currentDir = rotate(currentDir)
				}
			}
		}
	}
	fmt.Println(total)
}

func q6p1() {
	lines := readInGrid()
	imax, jmax := len(lines)-1, len(lines[0])-1
	i, j := locateGuard(lines)
	currentDir := lines[i][j]

	for i > 0 && j > 0 && i < imax && j < jmax {
		ii, jj := direction(currentDir)
		if checkNextSquare(lines, i, j, ii, jj) {
			lines[i][j] = been
			i += ii
			j += jj
		} else {
			currentDir = rotate(currentDir)
		}
	}

	total := 1
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			char := lines[i][j]
			if char == been {
				total++
			}
		}
	}
	fmt.Println(total)
}

func checkNextSquare(lines [][]byte, i int, j int, ii int, jj int) bool {
	nextSquare := lines[i+ii][j+jj]
	return nextSquare != block
}

func locateGuard(lines [][]byte) (int, int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			char := lines[i][j]
			if isGuard(char) {
				return i, j
			}
		}
	}
	panic("locateGuard")
}

func isGuard(char byte) bool {
	return char == up ||
		char == down ||
		char == left ||
		char == right
}

func rotate(char byte) byte {
	switch char {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	panic("rotate")
}

func direction(char byte) (int, int) {
	switch char {
	case up:
		return -1, 0
	case down:
		return 1, 0
	case left:
		return 0, -1
	case right:
		return 0, 1
	}
	panic("direction")
}

func readInGrid() [][]byte {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	return lines
}
