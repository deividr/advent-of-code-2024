package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile2() [][]int {
	var reports [][]int

	file, err := os.Open("day2/input.txt")

	if err != nil {
		panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		var lineInt []int

		for _, value := range line {
			intNumber, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Erro ao converter para int")
				break
			}
			lineInt = append(lineInt, intNumber)
		}

		reports = append(reports, lineInt)
	}

	return reports
}

func main() {
	reports := readFile2()
	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	safes := 0

	for _, values := range reports {
		if consist(values) {
			safes += 1
		}
	}

	fmt.Println(safes)
}

func part2(reports [][]int) {
	safes := 0

	for _, values := range reports {
		if consist(values) {
			safes += 1
			continue
		}

		for i := range values {
			newValues := make([]int, 0, len(values)-1)
			newValues = append(newValues, values[:i]...)
			newValues = append(newValues, values[i+1:]...)

			if consist(newValues) {
				safes += 1
				break
			}
		}
	}

	fmt.Println(safes)
}

func consist(values []int) bool {
	flux := ""
	safe := true

	for i, value := range values[1:] {
		if value == values[i] {
			safe = false
			break
		}

		if flux == "" {
			if value > values[i] {
				flux = "increasing"
			}

			if value < values[i] {
				flux = "decreasing"
			}
		}

		if flux == "decreasing" && value > values[i] {
			safe = false
			break
		}

		if flux == "increasing" && value < values[i] {
			safe = false
			break
		}

		diff := value - values[i]

		if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			safe = false
			break
		}
	}

	return safe
}
