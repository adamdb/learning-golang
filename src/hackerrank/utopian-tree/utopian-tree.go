package main

import (
  "fmt"
  "strconv"
  "os"
  "math"
)

func getTreeHeight(cycles int) int {
  h := 1
  
  for i := 0; i < cycles; i++ {
    if math.Mod(float64(i), 2) == 0 {
      h = h * 2
    } else {
      h = h + 1
    }
  }

  return h
}

func main() {
  T, _ := strconv.Atoi(os.Args[1])

  for i := 2; i < (T + 2); i++ {
    N, _ := strconv.Atoi(os.Args[i])
    h := getTreeHeight(N)
    
    fmt.Println(h)
  }
}
