package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

func readFile() ([]int, []int) {
  var (
    left  []int
    right []int
  )

  file, err := os.Open("day1/input.txt")
  if err != nil {
    panic(fmt.Sprintf("failed to open file\n\nerr:\n%v\n", err))
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := strings.Split(strings.Trim(scanner.Text(), " "), " ")

    leftInt, err := strconv.Atoi(line[0])
    if err != nil {
      panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
    }

    rightInt, err := strconv.Atoi(line[len(line)-1])
    if err != nil {
      panic(fmt.Sprintf("failed to convert left string to int\n\nerr:\n%v\n", err))
    }

    left = append(left, leftInt)
    right = append(right, rightInt)
  }

  sort.Slice(left, func(a, b int) bool {
    return left[a] < left[b]
  })

  sort.Slice(right, func(a, b int) bool {
    return right[a] < right[b]
  })

  return left, right
}

func main() {
  left, right := readFile()
  part1(left, right)
  part2(left, right)
}

func part1(left, right []int) {
  value := 0

  for i := range left {
    value += absInt(left[i] - right[i])
  }

  fmt.Println(value)
}

func part2(left, right []int) {
  counts := make(map[int]int)
  value := 0

  for _, num := range right {
    counts[num]++
  }

  for _, num := range left {
    value += counts[num] * num
  }

  fmt.Println(value)
}

func absInt(num int) int {
  if num < 0 {
    return -num
  }
  return num
}
