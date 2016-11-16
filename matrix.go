package main

import (
  "fmt"
  matrix "github.com/skelterjohn/go.matrix"
)

func main() {
  // Slice of Slice
  dary := make([][]int, 4)
  for i := range dary {
    dary[i] = make([]int, 3)
  }
  dary[0][1] = 7
  fmt.Println(dary) // => "[[0 7 0] [0 0 0] [0 0 0] [0 0 0]]"

  // Matrix
  dmat := matrix.Zeros(4, 3)
  dmat.Set(0, 1, 7)
  fmt.Println(dmat)
}
