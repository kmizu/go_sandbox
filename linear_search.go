package main

import "fmt"

func linear_search(a []interface{} , x interface{}) int {
  for i := 0; i < len(a); i += 1 {
    if a[i] == x {
      return i
    }
  }
  return -1
}

func main() {
  langs := []interface{}{"Scala", "Haskell", "Golang"}
  index := linear_search(langs, "Haskell")
  fmt.Println(index)
}
