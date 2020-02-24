package main

import (
  "fmt"
  "time"
)

const N int64 = 2000000

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

func pow(a, b, p int64) int64 {
  if b == 0 {
    return 1
  } else if b%2 == 0 {
    return pow(a*a%p, b/2, p)
  } else {
    return (a * pow(a*a%p, b/2, p)) % p
  }
}

func roots(p int64) []int64 {
  if p%3 != 1 {return []int64{p-1}}
  for a:=int64(2); a<p; a++{
    x := pow(a, (p-1)/3, p)
    if x!=1 {
      return []int64{p-x, p-(x*x%p), p-1}
    }
  }
  return []int64{}
}


func main() {
  start := time.Now()

  //for _, p := range primes_below(100) {
  //  fmt.Println(p)
  //}

  //fmt.Println(pow(10201, 281692, 97))
  //fmt.Println(roots(97))
  //fmt.Println(roots(101))

  lpf := make([]int64, N+1)
  for k:=int64(0) ; k < N+1 ; k++ {
    lpf[k] = 1+k*k*k
  }
  for _, p := range primes_below(N+3) {
    for _, s := range roots(p) {
      for v:=s ; v < N+1 ; v+=p {
        for (lpf[v]%p==0) && (lpf[v]>p) {
          lpf[v] /= p
        }
      }
    }
  }

  //for i:=int64(0) ; i < N+1 ; i++ {
  // fmt.Printf("lpf(%v", i)
  // fmt.Printf(") = %v\n", lpf[i])
  //}

  ans := int64(0)
  for _, p := range lpf {
    ans+=p
  }
  ans = ans - (N+1)

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))

}
