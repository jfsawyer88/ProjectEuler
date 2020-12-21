package main

import (
  "math/big"
  "fmt"
  "time"
)

const N = 2020
const m = 10000000000000000

func min(a, b int64) int64 {
  if a < b {
    return a
    } else {
      return b
    }
  }

func zeroes(l int64) []int64 {
    res := make([]int64, 0)
    for i := int64(0); i < l; i++ {
      res = append(res, 0)
    }
    return res
  }

func ones(l int64) [][]int64 {
  res := make([][]int64, 1)
  for i := int64(0); i < l; i++ {
    res[0] = append(res[0], 1)
  }
  return res
}

func partitions(n int64) [][]int64 {
  // retunrs the partitions of n
  var ans [][]int64
  for k := int64(1); k < n+1; k++ {
    ans = append(ans, partitions_sub(n, k)...)
  }
  return ans
}

func partitions_sub(n int64, k int64) [][]int64 {
  // returns the partitions of n with largest summand=k
  if k==1 { return ones(n) }
  if k==n { return [][]int64{{n}} }
  var ans [][]int64
  for i := int64(1); i < 1+min(k, n-k); i++ {
    for _, p := range partitions_sub(n-k, i) {
      ans = append(ans, append(p, k))
    }
  }
  return ans
}

func valid_partitions(n, k int64) [][]int64 {
  // partitions of k with count of summands < n, with k appended
  var ans [][]int64
  for _, p := range partitions(k) {
    if int64(len(p)) < n {
      p = append(p, k)
      p = append(zeroes(n-int64(len(p))), p...)
      ans = append(ans, p)
    }
  }
  return ans
}

func factorial(n int64) map[int64]int64 {
  // returns the factorial of n in factorized form
  // represented as a map of primes -> exponents
  ans := make(map[int64]int64)
  for _, p := range primes_below(n+1) {
    ans[p] = factorial_multiplicity(n, p)
  }
  return ans
}

func factorial_multiplicity(n, p int64) int64 {
  e := int64(0)
  pp := p
  for n/pp > 0 {
    e += n/pp
    pp *= p
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

func overcounting_factor(vp []int64) map[int64]int64{
  // returns the overcounting factor of a valid partition vp
  digit_count := make(map[int64]int64)
  for _, d := range vp {
    digit_count[d] += 1
  }
  ans := make(map[int64]int64)
  for _, c := range digit_count {
    ans = multiply(ans, factorial(c))
  }
  return ans
}

func multiply(a, b map[int64]int64) map[int64]int64 {
  // returns the product of a and b
  // inputs are fully factorized, represented as a map of primes -> exponents

  // copy of a
  res := make(map[int64]int64)
  for p, e := range a {
    res[p] = e
  }

  // add in prime factors of b into copy of a
  for p, e := range b {
    res[p] += e
  }
  return res
}

func divide(a, b map[int64]int64) map[int64]int64 {
  // returns quotient of a by b
  // inputs are fully factorized, represented as a map of primes -> exponents
  // assumpes b divides a

  // copy of a
  res := make(map[int64]int64)
  for p, e := range a {
    res[p] = e
  }

  // remove factors of b from factors of a
  for p, e := range b {
    res[p] -= e
  }
  return res
}

func factorize(n int64) map[int64]int64 {
  ans := make(map[int64]int64)

  for p := int64(2) ; p*p <= n ; p++ {
    for 0 == n%p {
      ans[p] += 1
      n/=p
    }
  }
  // if n is not fully reduced to 1 then the remaining value is prime itself
  if 1 < n {
    ans[n] = 1
  }
  return ans
}

func mod_pow(a, b, p int64) int64 {
  if b == 0 {
    return 1
  } else if b%2 == 0 {
    return mod_pow(a*a%p, b/2, p)
  } else {
    return (a * mod_pow(a*a%p, b/2, p)) % p
  }
}

func evaluate(a map[int64]int64, q int64) *big.Int {
  // evaluate factorized a mod q
  ans := big.NewInt(int64(1))
  pe := big.NewInt(int64(0))
  for p, e := range a {
    //ans *= mod_pow(p, e, q)
    pe.Exp(big.NewInt(p), big.NewInt(e), big.NewInt(q))
    ans.Mul(ans, pe)
  }
  return ans
}

func S_sub(n, k, q int64) *big.Int {
  sum := big.NewInt(int64(0))
  full := multiply(factorize(2*k), factorial(n-1))
  for _, p := range valid_partitions(n, k) {
    sum.Add(sum, evaluate(divide(full, overcounting_factor(p)), q))
    sum.Mod(sum, big.NewInt(q))
  }
  return sum
}

func S(n, q int64) *big.Int {
  ans := big.NewInt(int64(0))
  for k := int64(1); k < 10; k++ {
    //ans += S_sub(n, k, q)
    ans.Add(ans, S_sub(n, k, q))
  }
  ones := big.NewInt(int64(0))
  for i := int64(0); i < n; i++ {
    //ones = (ones + mod_pow(10, i, q)) % q
    ones.Add(ones, big.NewInt(mod_pow(10, i, q)))
    ones.Mod(ones, big.NewInt(q))
  }
  //return (ans * ones) % q
  ans.Mul(ans, ones)
  ans.Mod(ans, big.NewInt(q))
  return ans
}

func main() {
  start := time.Now()
  ans := S(N, m)
  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
