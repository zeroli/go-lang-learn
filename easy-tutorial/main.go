package main

import (
  "fmt"
  //"os"
  "math/rand"
  "sort"
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

  LearnStruct()
  LearnArray()
  LearnSlice()
  LearnMap()
  LearnInterface()
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
  {
    scores := []int{1, 2, 3}
    fmt.Println(scores)
    ModifyArray(scores)
    fmt.Println(scores) // modified in place
  }
}

func LearnSlice() {
  {
    scores := []int{1, 2, 3}
    fmt.Println("make slice with literals")
    for index, value := range scores {
      fmt.Printf("scores[%d]=%d\n", index, value)
    }
  }
  {
    scores := make([]int, 0)
    scores = append(scores, 3)
    fmt.Println("make slice:")
    for index, value := range scores {
      fmt.Printf("scores[%d]=%d\n", index, value)
    }
  }
  {
    scores := make([]int, 100)
    for i := 0; i < 100; i++ {
      scores[i] = int(rand.Int31n(1000))
    }
    sort.Ints(scores)

    fmt.Println(scores)
    worst := make([]int, 5)
    // only copy the first 2 elements to worst[2:4]
    copy(worst[2:4], scores[:5])
    fmt.Println(worst)
    // copying, once one (src/dst) ends, stop copying
    copy(worst, scores[:1])
    fmt.Println(worst)
  }

}

func LearnMap() {
  {
    lookup := make(map[string]int)
    lookup["goku"] = 9000
    power, exists := lookup["vegeta"]

    fmt.Println(power, exists)
    fmt.Println("size of map:", len(lookup))
    delete(lookup, "goku")
    delete(lookup, "xxx")
    fmt.Println("size of map:", len(lookup))
  }
  {
    lookup := map[string]int {
      "goku" : 9000,
      "xxx" : 1,
    }
    fmt.Println(lookup)
  }
  {
    lookup := map[string]int {
      "a" : 10,
      "b" : 20,
      "c" : 30,
    }
    for k, v := range(lookup) {
      fmt.Println(k, v)
    }
  }
}

type Logger interface {
  Log(message string)
}

type ConsoleLogger struct {
}

func (logger ConsoleLogger) Log(message string) {
  fmt.Println(message)
}

func logMessage(logger Logger, message string) {
  logger.Log(message)
}

func LearnInterface() {
  {
    logger := &ConsoleLogger{}
    logMessage(logger, "hello interface to log message")
  }
}

// passing array as a reference
func ModifyArray(a []int) {
  a[0] = 10
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
