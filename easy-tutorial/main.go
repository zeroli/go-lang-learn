package main

import (
  "fmt"
  "os"
)

func log(message string) {

}

// (int) or int as return, are ok
func add(a int, b int) (int) {
  return 1
}

func power(name string) (int, bool) {
  return 1, false
}

func main() {
  power := 9000
  // compiler error
  //power := 9001
  x, power := 1, 8000  // ok, since x is new
  fmt.Println("x = ", x)
  fmt.Printf("it's over %d\n", power)
  if len(os.Args) != 2 {
    os.Exit(1);
  }
}
