package main

import (
  "fmt"
  "time"
)

const LIM int = 4000000

func main() {

  start := time.Now()

  var sum int = 0

  fib := [2]int {0, 1}
  for fib[1] < LIM {
    if fib[1] % 2 == 0 {
      sum += fib[1]
    }
    fib = [2] int{fib[1], fib[0] + fib[1]}
  }

  fmt.Printf("Result is: %v\n", sum)
  fmt.Printf("Timing: %v", time.Since(start))

}
