package main

import (
  "fmt"
  "time"
)



func main() {

  start := time.Now()
  //var N in = 600851475143
  var N int = 377284927380394081

  for N%2 == 0{
    N /= 2
  }

  var lpf int = 3
  for N > lpf*lpf {
    for N%lpf == 0 {
      N /= lpf
    }
    lpf += 2
  }

  if N > lpf {
    lpf = N
  }

  fmt.Printf("Result is: %v\n", N)
  fmt.Printf("Timing: %v", time.Since(start))

}
