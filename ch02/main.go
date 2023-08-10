package main

import (
  "fmt"
  "strconv"
  "strings"
)


func main() {
  // multiple var declarations could group in ()
  // and without type suffix
  {  
    var (
      i = 10
      j = 30.2
    )
    fmt.Println(i, j)
  }
  {
    var f32 float32 = 2.2
    var f64 float64 = 10.2345
    fmt.Println("f32 is", f32, ", f64 is", f64)
  }

  // string
  {
    var s1 = "hello"
    var s2 = "world"
    fmt.Println(s1 + s2)
  }

  // uninitialized var will be auto-initialized to ZERO value
  {
    var zi int  // zi = 0
    var zf float64 // zf = 0.0
    var zb bool // zb = false
    var zs string // sz = ""
    fmt.Println(zi, zf, zb, zs)
  }

  // shorter var declaration without 'var'
  {
    i := 10
    bf:= false
    s1:= "hello"
    fmt.Println(i, bf, s1)
  }
  // pointer variable
  {
    i := 10
    p := &i
    fmt.Println(*p)
  }
  // const variable, cannot be modified at runtime
  // only primitive types can be const
  {
    // cannot use `const i := 10`
    const i int = 10
    // i = 20  // compile error
    fmt.Println(i)
    const (one = 1; two = 2; three = 3)
    fmt.Println(one, two, three)

    // use `iota` constant generator
    const (f1 = iota+1
      f2
      f3
      f4
    )
    fmt.Println(f1, f2, f3, f4)  // 1, 2, 3, 4
  }
  // type conversion
  {
    i := 10
    i2s := strconv.Itoa(i)
    s2i,_ := strconv.Atoi(i2s)
    fmt.Println(i, i2s, s2i)

    f := float32(i)
    fmt.Println(f)
  }
  // string pacakge
  {
    fmt.Println(strings.HasPrefix("hello", "h"))
    fmt.Println(strings.Index("hello", "o"))
  }
}
