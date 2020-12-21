package main

import (
  "math/big"
  "fmt"
  "time"
)

const N = 1000000000
const P = 100

func intLog(n, b *big.Int) *big.Int {
  // integer log of n base b
  // aka, the smallst e such that b^e<=n
  e := big.NewInt(int64(0))
  for bp := big.NewInt(int64(1)) ; n.Cmp(big.NewInt(0).Mul(bp, b)) > -1 ; bp.Mul(bp, b) {
    e.Add(e, big.NewInt(int64(1)))
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

func ham(N *big.Int, primes []int64) *big.Int {
  // returns count of "hamming" numbers not exceeding N
  // which only have prime factors in primes
  // nb: primes must be sorted

  // if only one prime left then only powers available, so take a log
  if 1 == len(primes) {
    return big.NewInt(0).Add(big.NewInt(int64(1)), intLog(N, big.NewInt(primes[0])))
  }
  // if largest prime to consider is too large then drop it
  if N.Cmp(big.NewInt(primes[len(primes)-1])) == -1 {
    return ham(N, primes[:len(primes)-1])
  }

  // recursively
  return big.NewInt(0).Add(ham(N, primes[:len(primes)-1]),ham(big.NewInt(0).Div(N, big.NewInt(primes[len(primes)-1])), primes))
}


func main() {
  start := time.Now()
  //ans := ham(big.NewInt(2), []int64{2})
  //ans := intLog(big.NewInt(2), big.NewInt(2))
  ans := ham(big.NewInt(37037037037037), []int64{2, 3, 5})
  //ans := ham(big.NewInt(5), []int64{2})
  //ans := ham(N, primes_below(P))
  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
