package main

import (
  "fmt"
  "time"
  //"math"
)

const N = 1000000

func collatz(n int) int {
  if n == 0 {
    return 1
  } else if 0 == n%2 {
    return n/2
  } else {
    return 1+3*n
  }
}

func main() {

  start := time.Now()

  memo := make(map[int] int)
  memo[1] = 1
  var s, v, m int
  var ok bool
  //for n:=N/2 ; n<N ; n++ {
  for n:=1 ; n<N ; n++ {
    s = 0
    m = n
    for v, ok = memo[m]; !ok ; v, ok = memo[m] {
      //fmt.Println(v, ok, n, m, s)
      m = collatz(m)
      s += 1
    }
    memo[n] = v + s
  }

  ans := 0
  max_chain := 0
  for n, chain := range memo {
    if max_chain < chain {
      ans = n
      max_chain = chain
    }
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))

}
