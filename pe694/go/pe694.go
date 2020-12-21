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

func ham(N int64, primes []int64) int64 {
  // returns count of "hamming" numbers not exceeding N
  // which only have prime factors in primes
  // nb: primes must be sorted

  // if only one prime left then only powers available, so take a log
  if 1 == len(primes) {
    return 1+intLog(N, primes[0])
  }
  // if largest prime to consider is too large then drop it
  if N < primes[len(primes)-1] {
    return ham(N, primes[:len(primes)-1])
  }

  // recursively
  return ham(N, primes[:len(primes)-1]) + ham(N/primes[len(primes)-1], primes)
}





func main() {
  start := time.Now()



  //fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
