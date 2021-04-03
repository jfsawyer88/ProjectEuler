package main

import (
  "math"
  "fmt"
  "time"
)

const N = 1E36

func nRoot(n int64, N float64) int64 {
  return int64(math.Pow(N, 1/float64(n)))
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

func ham(N int64, primes []int64) int64 {
  if 0 == len(primes) { return 1 }
  if 0 == N { return 0 }
  return ham(N, primes[:len(primes)-1]) + ham(N/primes[len(primes)-1], primes)
}


func main() {
  start := time.Now()
  //ans := ham(100000000, []int64{2, 3, 5})
  //ans := ham(1000000000000000000/133049351085651000, primes_below(16))
  //ans := ham(4, []int64{2, 3})
  //ans := ham(1000000000000000000/(9261000), []int64{2, 3, 5, 7})
  //ans := ham2(1000000000000000000/(9261000), 1, []int64{2, 3, 5, 7})
  //ans := ham(N, primes_below(P))
  //ans := int64(math.Pow(N, 1.0/5))
  //primes_below(1000000000)
  //primes_below(15848931)
  //ans := nRoot(13, N)
  //fmt.Printf("Value is %v\n", N)
  //fmt.Printf("Type is is %T\n", N)
  //fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
