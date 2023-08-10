package main

import (
  "fmt"
)

func main() {
  // the if/else block must have '{}' to wrap
  // and '{' must follow condition on same line
  {
    i := 6
    if i > 10 {
      fmt.Println("i > 10")
    } else if i > 5 && i <= 10 {
      fmt.Println("5 < i <= 10")
    } else {
      fmt.Println("i <= 5")
    }
  }
  // if check could have simple expression separated with condition expr
  {
    if i := 6; i > 10 {
      fmt.Println("i > 10")
    } else if i > 5 && i <= 10 {
      fmt.Println("5 < i <= 10")
    } else {
      fmt.Println("i <= 5")
    }
  }
  // switch case for more if/else check
  // no 'break' for each case, auto break once match and executed
  // 'fallthrough' to fall through next case explicitely
  {
    // the "{" after switch could be on new line    
    switch i := 6
    {
    // switch i := 6; {   // note: there is one ";"
    case i > 10:
      fmt.Println("i > 10")
    case i > 5 && i <= 10:
      fmt.Println("5 < i <= 10")
    default:
      fmt.Println("i <= 5")
    }
  }
  // for loop, go doesn't have "while loop"
  {
    sum := 0
    for i := 1; i <= 100; i++ {
      sum += i
    }
    fmt.Println(sum)
  }
  // go doesn't have 'while loop', use shorter 'for loop' instead
  {
    sum := 0
    i := 1
    for i <= 100 {
      sum += i
      i++
    }
    fmt.Println(sum)
  }
  // continue/break
  {
    sum := 0
    i := 1
    for {  // infinite loop
      sum += i
      i++
      if i > 100 {
        break
      }
    }
    fmt.Println(sum)
  }

}

