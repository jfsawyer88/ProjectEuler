package main

import (
  "fmt"
  "time"
)

//const N = 1000000000
//const P = 100
const N = 100000000000000
const P = 100


func intLog(n, b int64) int64 {
  // integer log of n base b
  // aka, the smallst e such that b^e<=n
  e := int64(0)
  for bp := int64(1) ; bp*b <= n ; bp=bp*b {
    e += 1
  }
  return e
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
  // returns count of "hamming" numbers not exceeding N
  // which only have prime factors in primes
  // nb: primes must be sorted

  // if only one prime left then only powers available, so take a log
  //if 0 == len(primes) { return 1 }
  if 1 == len(primes) {
    return 1+intLog(N, primes[0])
  }

  // if largest priem to consider is too large then drop it
  //if 0 == N { return 0 }
  if N < primes[len(primes)-1] {
    return ham(N, primes[:len(primes)-1])
  }

  // recursively
  return ham(N, primes[:len(primes)-1]) + ham(N/primes[len(primes)-1], primes)
}

func ham2(lim int64, N int64, primes []int64) int64 {
  //ans := int64(1)
  ans := N // <- each of these is a hamming number
  for i := 0; (i < len(primes) && (N*primes[i]) <= lim) ; i++ {
    ans += ham2(lim, N*primes[i], primes[i:])
  }
  return ans
}

func main() {
  start := time.Now()
  //ans := ham(100000000, []int64{2, 3, 5})
  //ans := ham(1000000000000000000/133049351085651000, primes_below(16))
  //ans := ham(4, []int64{2, 3})
  //ans := ham(1000000000000000000/(9261000), []int64{2, 3, 5, 7})
  //ans := ham2(1000000000000000000/(9261000), 1, []int64{2, 3, 5, 7})
  //ans := ham(N, primes_below(P))
  ans := ham2(N, 1, primes_below(P))
  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
