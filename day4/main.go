package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() []string {
	var lines []string

	file, err := os.Open("day4/input.txt")

	if err != nil {
		panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	lines := readFile()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for j, line := range lines {
		for i := 0; i < len(line); i++ {
			if i < len(line)-3 && (line[i:i+4] == "XMAS" || line[i:i+4] == "SAMX") {
				total += 1
			}

			if j > len(lines)-4 || (line[i] != 'X' && line[i] != 'S') {
				continue
			}

			// vertical
			horizontal := string([]byte{line[i], lines[j+1][i], lines[j+2][i], lines[j+3][i]})
			if horizontal == "XMAS" || horizontal == "SAMX" {
				total += 1
			}

			// diagonal - from right to left
			if i > 2 {
				diagonal := string([]byte{line[i], lines[j+1][i-1], lines[j+2][i-2], lines[j+3][i-3]})
				if diagonal == "XMAS" || diagonal == "SAMX" {
					total += 1
				}
			}

			// diagonal - from left to right
			if i < len(line)-3 {
				diagonal := string([]byte{line[i], lines[j+1][i+1], lines[j+2][i+2], lines[j+3][i+3]})
				if diagonal == "XMAS" || diagonal == "SAMX" {
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func part2(lines []string) {
	total := 0
	for j, line := range lines[:len(lines)-2] {
		for i := 0; i < len(line)-2; i++ {
			if line[i] != 'M' && line[i] != 'S' {
				continue
			}

			de := string([]byte{line[i], lines[j+1][i+1], lines[j+2][i+2]})
			dd := string([]byte{line[i+2], lines[j+1][i+1], lines[j+2][i]})

			if (de == "MAS" || de == "SAM") && (dd == "MAS" || dd == "SAM") {
				total += 1
			}
		}
	}

	fmt.Println(total)
}
