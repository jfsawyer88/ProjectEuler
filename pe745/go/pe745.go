package main

import (
  "fmt"
  "time"
)

const LIM = 100000000000000
const m  = 1000000007

func abs(n int64) int64 {
  if n < 0 {
    return -n
  } else {
    return n
  }
}

func iSqrt(n int64) int64 {
  var x int64 = n
  xk := (x+n/x)
  xk1 := (xk-n/xk)/2
  for (abs(xk1-xk)) > 1 {
    xk = xk1
    xk1 = (xk+n/xk)/2
  }
  return xk1
}

func primes_below(N int64) []int64 {
  is_prime := make([]bool, N)

  for i := int64(2); i < N; i++ {
    is_prime[i] = true
  }

  p := int64(2)
  for {
    p2 := p * p
    if p2 >= N {
      break
    }
    // second allowed optimization:  inner loop starts at sqr(p)
    for i := p2; i < N; i += p {
      is_prime[i] = false // it's a composite
    }
    // scan to get next prime for outer loop
    for {
      p++
      if is_prime[p] {
        break
      }
    }
  }
  cnt := int64(0)
  for p:=int64(0) ; p < N ; p++ {
    if is_prime[p] {
      cnt++
    }
  }
  primes := make([]int64, cnt)
  i := int64(0)
  for p:=int64(0) ; p < int64(len(is_prime)) ; p++ {
    if is_prime[p] {
      primes[i] = p
      i++
    }
  }
  return primes
}

func prd(A []int64) int64 {
  ans := int64(1)
  for _, a := range A {
    ans *= a
  }
  return ans
}

func sqr(n int64) int64 {
  return n*n
}

func square_free_count(LIM int64) int64 {
  return square_free_count_sub(1, 1, primes_below(1+iSqrt(LIM)), LIM)
}

func square_free_count_sub(currProd int64, inclexcl int64, remaining []int64, LIM int64) int64 {
  //fmt.Printf("currProd=%v, inclexcl=%v, LIM=%v\n", currProd, inclexcl, LIM)
  ans := inclexcl*LIM / sqr(currProd)
  for i, p := range remaining {
    if LIM/sqr(p) < sqr(currProd) { break }
    ans += square_free_count_sub(currProd*p, inclexcl*-1, remaining[i+1:], LIM)
  }
  return ans
}


func main() {
  start := time.Now()

  ans := int64(0)
  for i := int64(1); i*i <= LIM; i++ {
    //fmt.Printf("i=%v, i*i=%v, newlim=%v, sqfree=%v\n", i, i*i, LIM/(i*i), square_free_count(LIM/(i*i)))
    ans = (ans+i*i*square_free_count(LIM/(i*i))) % m
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
