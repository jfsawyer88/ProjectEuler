package main

import (
  "fmt"
  "time"
)

//const N = 10000
const N = 1000000000000

func pow(a, b int64) int64 {
  if b == 0 {
    return 1
  } else if b%2 == 0 {
    return pow(a*a, b/2)
  } else {
    return a * pow(a*a, b/2)
  }
}

func prepend(n int64, A []int64) []int64 {
  return append([]int64{n}, A...)
}

func integer_to_digits(n, b int64) []int64 {
  N := make([]int64, 0)
  for 0 < n {
    N = prepend(n%b, N)
    n /= b
  }
  return N
}

func digits_to_integer(N []int64, b int64) int64 {
  n := int64(0)
  for i, d := range N {
    n += d * pow(b, int64(len(N)-1-i))
  }
  return n
}

func possible_sums(A []int64) []int64 {
  if 0 == len(A) { return []int64{0} }
  if 1 == len(A) { return A }
  res := make([]int64, 0)
  for i := 1; i-1 < len(A); i++ {
    for _, A_sub := range possible_sums(A[i:]) {
      res = append(res, digits_to_integer(A[:i], 10)+A_sub)
    }
  }
  return res
}

// original solution: slightly optimized by ensuring
// at least one addend is large enough to meet the sum requirement
func is_sqrtS_number(sqrt_n int64) bool {
  // checks if n is an S-number
  // NB: input is the square root of n
  n := sqrt_n*sqrt_n
  N := integer_to_digits(n, 10)

  nd := len(integer_to_digits(sqrt_n, 10))
  for i := 0; i+nd-1 < len(N); i++ {
    A := N[:i]
    B := N[i:i+nd]
    C := N[i+nd:]

    b := digits_to_integer(B, 10)
    if (b > sqrt_n) || (digits_to_integer(A, 10)+b+digits_to_integer(C, 10)<sqrt_n) {
      continue
    }
    for _, a := range possible_sums(A) {
      if (a + b > sqrt_n) || (a + b + digits_to_integer(C, 10) < sqrt_n) { continue }
      for _, c := range possible_sums(C) {
        if a+b+c == sqrt_n {
          //fmt.Printf("%v: sqrt(%v) = %v+%v+%v\n", sqrt_n, n, a, b, c)
          return true
        }
      }
    }
  }

  nd -= 1
  for i := 0; i+nd-1 < len(N); i++ {
    A := N[:i]
    B := N[i:i+nd]
    C := N[i+nd:]

    b := digits_to_integer(B, 10)
    if (b > sqrt_n) || (digits_to_integer(A, 10) + b + digits_to_integer(C, 10)<sqrt_n) {
      continue
    }
    for _, a := range possible_sums(A) {
      if (a + b > sqrt_n) || (a + b + digits_to_integer(C, 10) < sqrt_n) { continue }
      for _, c := range possible_sums(C) {
        if a+b+c == sqrt_n {
          //fmt.Printf("%v: sqrt(%v) = %v+%v+%v\n", sqrt_n, n, a, b, c)
          return true
        }
      }
    }
  }

  return false
}

// simpler code that uses possible_sums more directly
// func is_sqrtS_number(sqrt_n int64) bool {
//     // checks if n is an S-number
//     // NB: input is the square root of n
//     n := sqrt_n*sqrt_n
//     N := integer_to_digits(n, 10)
//     for _, s := range possible_sums(N) {
//       if sqrt_n == s {
//         //fmt.Printf("sqrt(%v)=%v=%v\n", n, sqrt_n, s)
//         return true
//       }
//     }
//     return false
//   }


func main() {
  start := time.Now()
  ans := int64(0)
  for n := int64(4); n*n <= N; n++ {
    if (0 == n%9) || (1 == n%9) { // optimization due to ecnerwala based on digit sum observation
      if is_sqrtS_number(n) {
        ans += n*n
      }
    }
  }
  //ans := possible_sums(integer_to_digits(8100, 10))
  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
