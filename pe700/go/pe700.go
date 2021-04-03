package main

import (
  "fmt"
  "time"
)

const N = 100000000
const a = 1504170715041707
const a_inv = 3451657199285664
const p = 4503599627370517

func egcd(a, b int64) []int64 {
  x := int64(0)
  y := int64(1)
  u := int64(1)
  v := int64(0)
  var q, r, m, n int64
  for a > 0 {
    q = b / a
    r = b % a

    m = x - u*q
    n = y - v*q

    b = a
    a = r
    x = u
    y = v
    u = m
    v = n
  }
  return []int64{b, x, y}
}

func mod_inv(a, p int64) int64 {
  g := egcd(a, p)
  ans := g[1] % p
  if ans < 0 {
    ans += p
  }
  return ans
}

func mod_div(a, b, p int64) int64 {
  return mod_product(a, mod_inv(b, p), p)
}

func mod_product(a, b, p int64) int64 {
  ans := int64(0)
  a %= p
  for b > 0 {
    if 1 == b % 2 { ans = (ans+a)%p }
    a = (a*2)%p
    b /= 2
  }
  return ans
}

// func mod_product(a, b, p int64) int64 {
//   if b == 0 {
//     return 1
//   } else if b%2 == 0 {
//     return mod_product(a+a%p, b/2, p)
//   } else {
//     return (a + mod_product(a+a%p, b/2, p)) % p
//   }
// }

func sum(A []int64) int64 {
  res := int64(0)
  for _, a := range A {
    res += a
  }
  return res
}

// todo: add other half (calculating position of all lower possible Eulercoins,
//                      sorting them by index and only keeping those < all prior)
func main() {
  start := time.Now()
  //ans := make([]int64, 0)

  ec := []int64{a}
  var s int64
  for n := int64(2); n < N; n++ {
    s = mod_product(a, n, p)
    if s < ec[len(ec)-1] {
      ec = append(ec, s)
      fmt.Printf("s(%v) = %v is an Eulercoin!\n", n, s)
    }
  }

  ans := sum(ec)
  // ans := mod_div(4046188430, a, p)
  // ans := mod_div(14578937, a, p)
  // ans := mod_inv(a, p)
  // for m := int64(14578937-1); m > 0; m-- {
  //   mod_product(m, a_inv, p)
  //   //fmt.Printf("s(%v) = %v may be an Eulercoin\n", mod_product(m, a_inv, p), m)
  // }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
