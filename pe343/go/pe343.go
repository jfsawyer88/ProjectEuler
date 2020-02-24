package main

import (
  "fmt"
  "time"
)

const N int64 = 1+2000000

func abs(n int64) int64 {
  if n<0 {
    return -n
    } else {
      return n
    }
  }

func iSqrt(n int64) int64 {
  var x int64 = n
  xk := (x+n/x)/2
  xk1 := (xk+n/xk)/2
  for abs(xk1-xk) > 1 {
    xk = xk1
    xk1 = (xk+n/xk)/2
  }
  return xk1
}

func min(a, b int64) int64 {
  if a < b {
    return a
    } else {
      return b
    }
  }

func max(a, b int64) int64 {
  if a > b {
    return a
    } else {
      return b
    }
  }

func full_div(n, p int64) int64 {
  for n%p == 0 {
    n/=p
  }
  return n
}

func main () {
  start := time.Now()

// initialize smallest prime factor (spf) array as range
  spf := make([] int64, N+1)
  for i:=int64(0) ; i < N+1 ; i++ {
    spf[i] = i
  }
  for p:=int64(2) ; p*p < N+1 ; p++ {
    if spf[p] == p {
      for n:=p*p ; n < N+1 ; n+=p {
        spf[n] = min(spf[n], p)
      }
    }
  }

  var n int64
  lpf := make([] int64, N+1)
  for i:=int64(2) ; i <= N ; i++ {
    n = i
    for spf[n]!=n {
      n /= spf[n]
    }
    lpf[i] = n
  }

  prime_cnt := int64(0)
  for i:=int64(2) ; i < N+1 ; i++ {
    if spf[i]==i {
      prime_cnt+=int64(1)
    }
  }

  primes := make([] int64, prime_cnt)
  prime_indx := int64(0);
  for i:=int64(2) ; i < N+1 ; i++ {
    if spf[i] == i {
      primes[prime_indx] = i
      prime_indx++
    }
  }

  //for i:=int64(0) ; i < N+1 ; i++ {
  // fmt.Printf("spf(%v", i)
  // fmt.Printf(") = %v\n", spf[i])
  //}

  //for i:=int64(0) ; i < N+1 ; i++ {
  // fmt.Printf("lpf(%v", i)
  // fmt.Printf(") = %v\n", lpf[i])
  //}

  //for i:=int64(0) ; i < N+1 ; i++ {
  // fmt.Printf("prime(%v", i)
  // fmt.Printf(") = %v\n", primes[i])
  //}

  res := int64(0)
  var a, b, lpfa, lpfb int64
  for k:=int64(1) ; k < N ; k++ {
    a = k + 1
    b = k*k - k + 1

    lpfa = lpf[a]
    lpfb = 1
    for _, p := range primes {
      if b<N+1 {break}
      if b%p == 0 {
        lpfb = p
        b = full_div(b, p)
      }
    }
    if b < N+1 {
      lpfb = max(lpfb, lpf[b])
      } else {
        lpfb = b
      }

      //fmt.Println(k, max(lpfa, lpfb))
      res += max(lpfa, lpfb) - 1

    }

    fmt.Printf("Result is %v\n", res)
    fmt.Printf("Timing: %v", time.Since(start))

  }
