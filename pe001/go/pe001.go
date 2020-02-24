package main

import (
  "fmt"
  "time"
  )

const N int = 10000000

func main() {

  start := time.Now()

  var sum int = 0

  for n := 0 ; n < N; n++ {
    if (n%3 == 0) || (n%5 == 0) {
      sum += n
    }
  }

  fmt.Printf("Result is: %v\n", sum)
  fmt.Printf("Timing: %v", time.Since(start))

}
