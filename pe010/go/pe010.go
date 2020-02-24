package main

import (
  "fmt"
  "time"
)

const N = 2000000

func prime_sieve(N int) []int {
  is_prime := make([]bool, N)

  for i := 2; i < N; i++ {
    is_prime[i] = true
  }

  p := 2
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
  cnt := 0
  for p:=0 ; p < N ; p++ {
    if is_prime[p] {
      cnt++
    }
  }
  primes := make([]int, cnt)
  i := 0
  for p:=0 ; p < len(is_prime) ; p++ {
    if is_prime[p] {
      primes[i] = p
      i++
    }
  }

  return primes
}

func main() {

  start := time.Now()

  sum := 0
  primes := prime_sieve(N)
  for i:=0 ; i < len(primes) ; i++ {
    sum += primes[i]
  }

  fmt.Printf("Result is %v\n", sum)
  fmt.Printf("Timing: %v", time.Since(start))

}
