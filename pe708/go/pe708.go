package main

import (
  "fmt"
  "time"
)

//const N int64 = 100000000000000
const N int64 = 2000000000

func abs(n int64) int64 {
  if n < 0 {
    return -n
  } else {
    return n
  }
}

func iSqrt(n int64) int64 {
  var x int64 = n
  xk := (x+n/x)/2
  xk1 := (xk-n/xk)/2
  for (abs(xk1-xk)) > 1 {
    xk = xk1
    xk1 = (xk+n/xk)/2
  }
  return xk1
}

func prime_pi(N int64) int64 {
  var r int64 = iSqrt(N)
  V1 := make([] int64, r)
  for i:=int64(0) ; i < r ; i++ {
    V1[i] = N/(i+1)
  }
  V2 := make([] int64, V1[r-1]-1)
  for i:=int64(0) ; i < V1[r-1]-1 ; i++ {
    V2[i] = (V1[r-1]-1)-i
  }
  V := append(V1, V2...)

  S := make(map[int64]int64)
  for i:=int64(0) ; i < int64(len(V)) ; i++ {
    //S[V[i]] = ((V[i]*(V[i]+1))/2)-1
    S[V[i]] = ((V[i]+1)/2)-1
  }
  for p:=int64(2) ; p < r+1 ; p++ {
    if S[p] > S[p-1] {
      sp := S[p-1]
      p2 := p*p
      for _, v := range V {
        if v < p2 {break}
        //S[v] -= p*(S[v/p] - sp)
        S[v] -= (S[v/p] - sp)
      }
    }
  }
  return int64(1)+S[N]
}

func main() {
  start := time.Now()

  res := prime_pi(N)

  fmt.Printf("Result is %v\n", res)
  fmt.Printf("Timing: %v", time.Since(start))
}
