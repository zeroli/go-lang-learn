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

type Person struct {
  Name string
}

type Saiyan1 struct {
  *Person
  Power int
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

  LearnStruct();
  LearnArray();
}

func LearnStruct() {
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
  {
    goku := NewSaiyan1("Goku", 100)
    fmt.Println(goku)  // &{Goku, 100}, which is a pointer
  }
  {
    goku := NewSaiyan2("Goku", 200)
    fmt.Println(goku)  // {Goku, 200}
  }
  {
    goku := new(Saiyan)
    // need manual assignment
    fmt.Println(goku)
  }
  {
    goku := &Saiyan{
      Name: "Goku",
      Power: 900,
    }
    fmt.Println(goku)
  }
  {
    goku := &Saiyan1{
      Person: &Person{"Goku"},
      Power: 100,
    }
    fmt.Println(goku)
  }
}

func LearnArray() {
  {
    var scores [10]int
    scores[0] = 1
    scores[1] = 2
    fmt.Println(scores)
  }
  {
    // 1,2,3,4,5 => compile error, out of bound
    // 1,2,3 => padding 0
    scores := [4]int{1, 2, 3, }
    fmt.Println(scores)
    fmt.Println("array.length=", len(scores))
    for index, value := range scores {
      fmt.Printf("scores[%d]=%d\n", index, value)
    }
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

func NewSaiyan1(name string, power int) *Saiyan {
  return &Saiyan{
    Name: name,
    Power: power,
  }
}

func NewSaiyan2(name string, power int) Saiyan {
  return Saiyan{
    Name: name,
    Power: power,
  }
}
