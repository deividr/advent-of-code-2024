package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)

func readFile() []string {
  var lines []string

  file, err := os.Open("day3/input.txt")

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
  for _, line := range lines {
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    matchesMul := re.FindAllString(line, -1)
    re2 := regexp.MustCompile(`\d{1,3}`)
    for _, value := range matchesMul {
      numbers := re2.FindAllString(value, -1)
      n1, _ := strconv.Atoi(numbers[0])
      n2, _ := strconv.Atoi(numbers[1])
      total += n1 * n2
    }
  }
  fmt.Println(total)
}

func part2(lines []string) {
  total := 0
  re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
  re2 := regexp.MustCompile(`\d{1,3}`)
  enable := true

  for _, line := range lines {
    matches := re.FindAllString(line, -1)
    var matchesMul []string

    for _, value := range matches {
      if value == "do()" {
        enable = true
        continue
      }

      if value == "don't()" {
        enable = false
        continue
      }

      if enable {
        matchesMul = append(matchesMul, value)
      }
    }

    for _, value := range matchesMul {
      numbers := re2.FindAllString(value, -1)
      n1, _ := strconv.Atoi(numbers[0])
      n2, _ := strconv.Atoi(numbers[1])
      total += n1 * n2
    }
  }

  fmt.Println(total)
}
