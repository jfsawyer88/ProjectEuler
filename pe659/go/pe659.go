package main

import (
  "fmt"
  "time"
  "math/rand"
)

const N int64 = 10000000

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

func mod_sqrt(n, p int64) int64 {
  //fmt.Println("mod_sqrt: START")
  //fmt.Println(n, p)
  Q := (p-1)/2
  S := int64(1)
  for Q%2 == 0{
    Q/=2
    S+=1
  }
  //fmt.Println(Q, S)

  z := int64(1)+rand.Int63n(p-1)
  for 1 == mod_pow(z, (p-1)/2, p) {
    z = int64(1)+rand.Int63n(p-1)
  }

  c := mod_pow(z, Q, p)
  R := mod_pow(n, (Q+1)/2, p)
  t := mod_pow(n, Q, p)
  M := S

  var b, i int64
  for t != 1 {
    //fmt.Println(c, R, t, M)
    i = 1
    for tp := t*t % p ; 1 != tp ; tp=tp*tp % p {
      i+=1
    }

    b = mod_pow(c, pow(2, M-i-1), p)
    R = R*b % p
    t = t*(b*b%p) % p
    c = b*b % p
    M = i
  }
  //fmt.Println("mod_sqrt: END\n")
  return min(R, p-R)
}

func mod_half(n, p int64) int64 {
  return (n * ((p+1)/2))%p
}

func mod_pow(a, b, p int64) int64 {
  ans := int64(1)
  a %= p
  for b > 0 {
    if 1 == b % 2 { ans = (ans*a)%p }
    a = (a*a)%p
    b /= 2
  }
  return ans
}

func pow(a, b int64) int64 {
  if b == 0 {
    return 1
  } else if b%2 == 0 {
    return pow(a*a, b/2)
  } else {
    return (a * pow(a*a, b/2))
  }
}

func full_div(n, p int64) int64 {
  for n%p == 0 {
    n/=p
  }
  return n
}

func main() {
  start := time.Now()

  P := make([]int64, N+1)
  R := make([]int64, N+1)
  for k := int64(0); k < N+1; k++ {
    P[k] = 1+4*k*k
    R[k] = 1+4*k*k
  }

  var np, i, b int64
  for _, p := range primes_below(iSqrt(4*N*N+1))[2:] {
    b = mod_half((p-1)/2, p)
    if 1 != mod_pow(b, (p-1)/2, p) {
      continue
    }
    np = mod_sqrt(b, p)
    P[np] = p
    R[np] = full_div(R[np], p)
    i = p
    for ; i < N+1 ; i+=p {
      P[i-np] = p
      R[i-np] = full_div(R[i-np], p)
      if i+np < N+1 {
        P[i+np] = p
        R[i+np] = full_div(R[i+np], p)
      }
    }
    if i-np < N+1 {
      P[i-np] = p
      R[i-np] = full_div(R[i-np], p)
    }
  }

  ans := int64(0)
  for k := int64(1) ; k < N+1 ; k++ {
    ans = (ans + max(P[k], R[k])) % 1000000000000000000
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))

}
