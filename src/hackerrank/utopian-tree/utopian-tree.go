//https://www.hackerrank.com/challenges/utopian-tree

package main

import (
  "fmt"
  "strconv"
  "os"
  "math"
  "bufio"
)

func getTreeHeight(cycles []int) {
  l := len(cycles)

  for i := 0; i < l; i++ {
    h := 1
    N := cycles[i]

    for j := 0; j < N; j++ {
      if math.Mod(float64(j), 2) == 0 {
        h = h * 2
      } else {
        h = h + 1
      }
    }

    fmt.Println(h)
  }
}

func main() {
  inputs := 0
  var cycles []int = nil
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    input, _ := strconv.Atoi(scanner.Text())

    if cycles == nil {
      cycles = make([]int, input)
    } else {
      cycles[inputs - 1] = input
    }

    inputs = inputs + 1

    if inputs > len(cycles) {
      break
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading input:", err)
  }

  getTreeHeight(cycles)
}
