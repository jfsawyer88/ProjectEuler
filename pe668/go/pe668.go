package main

import (
  "fmt"
  "time"
  "math"
)

const N int = 10000000000

func abs(n int) int {
  if n < 0 {
    return -n
  } else {
    return n
  }
}

func iLog(n int, b int) int {
  return int(math.Log(float64(n))/math.Log(float64(b)))
}

func iSqrt(n int) int {
  var x int = n
  xk := (x+n/x)
  xk1 := (xk-n/xk)/2
  for (abs(xk1-xk)) > 1 {
    xk = xk1
    xk1 = (xk+n/xk)/2
  }
  return xk1
}

func bin(Arr []int, n int) int {
  l, r, mid:= 0, len(Arr), len(Arr)/2
  if n > Arr[len(Arr)-1] {return r}
  for l <= r {
    mid = (l+r)/2
    if Arr[mid] < n {
      l = mid + 1
    } else if Arr[mid] > n{
      r = mid - 1
    } else {
      return mid+1
    }
  }
  return l
}

func two_way_merge_sort(A []int, B []int) []int {
  output := make([]int, len(A)+len(B))
  i, j, k := 0, 0, 0
  for (i < len(A)) && (j < len(B)) {
    if A[i] < B[j] {
      output[k] = A[i]
      i++
      k++
    } else {
      output[k] = B[j]
      j++
      k++
    }
  }
  for i < len(A) {
    output[k] = A[i]
    i++
    k++
  }
  for j < len(B) {
    output[k] = B[j]
    j++
    k++
  }
  return output
}

func primes_below(N int) []int {
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

func prime_powers_below(N int, p int) []int {
  powers := make([]int, 1+iLog(N, p))
  powers[0] = 1
  for i:=1 ; i < len(powers) ; i++ {
    powers[i] = powers[i-1]*p
  }
  return powers
}

func slice_below(A []int, n int) []int {
  outLen := bin(A, n)
  out := make([]int, outLen)
  for i:=0 ; i < outLen ; i++ {
    out[i] = A[i]
  }
  return out
}

func slice_scale(A []int, s int) []int {
  out := make([]int, len(A))
  for i:=0 ; i<len(out) ; i++ {
    out[i] = A[i]*s
  }
  return out
}

func main() {
  start := time.Now()

  psmooth := prime_powers_below(N, 2)
  ans := len(psmooth)-2

  for _, p := range primes_below(iSqrt(N))[1:] {
    //powers:= prime_powers_below(N/p, p)[1:]
    next_smooth := slice_below(psmooth, N/p)
    for _, ppow := range prime_powers_below(N/p, p)[1:] {
      next_smooth = two_way_merge_sort(next_smooth, slice_scale(slice_below(psmooth, N/(p*ppow)), ppow))
    }
    ans += len(next_smooth) - p
    psmooth = next_smooth
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
