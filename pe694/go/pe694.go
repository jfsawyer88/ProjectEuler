package main

import (
  "fmt"
  "time"
)

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

func ham2(lim int64, N int64, primes []int64) int64 {
  ans := lim/N
  for i := 0; i < len(primes) && N <= lim/primes[i] ; i++ {
    ans += ham2(lim, N*primes[i], primes[i:])
  }
  return ans
}

func pe694(N int64) int64 {
  // returns the sum of squarefree numbers below N
  return pe694_sub(N, make([]int64, 0), primes_below(N))
}

func pe694_sub(N int64, A, primes []int64) int64 {
  a := int64(1)
  for _, p := range A {
    a *= p
  }
  ans := ham2(N*N*N, a*a*a, A)
  for i := 0; (i < len(primes)) && (a <= N/primes[i]); i++ {
    ans += pe694_sub(N, append(A, primes[i]), primes[i+1:])
  }
  return ans
}

func main() {
  start := time.Now()
  ans := pe694(1000000)
  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
