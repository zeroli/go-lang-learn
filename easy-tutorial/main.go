package main

import (
  "fmt"
  //"os"
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

type Saiyan struct {
  Name string
  Power int
}

func main() {
  power := 9000
  // compiler error
  //power := 9001
  x, power := 1, 8000  // ok, since x is new
  fmt.Println("x = ", x)
  fmt.Printf("it's over %d\n", power)

  {
    goku := Saiyan{
      Name: "Goku",
      Power: 9000,
    }
    fmt.Println(goku)
    Super1(goku)
    fmt.Println(goku)  // {Goku, 9000}
  }
  {
    goku := Saiyan{}
    goku.Name = "Goku"
    goku.Power = 900
    fmt.Println(goku)
    Super2(&goku)
    fmt.Println(goku)  // {Goku, 1000}
  }
  {
    goku := Saiyan{Name: "Goku"}
    goku.Power = 100
    fmt.Println(goku)
    (&goku).Super();
    fmt.Println(goku)  // {Goku, 300}
  }
}

func Super1(s Saiyan) {
  s.Power += 100
}

func Super2(s *Saiyan) {
  s.Power += 100
}

func (s *Saiyan) Super() {
  s.Power += 200
}
