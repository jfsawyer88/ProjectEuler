package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()
  ans := int64(0)
  a := int64(1504170715041707)
  p := int64(4503599627370517)
  var new_a, new_p int64
  for a > 0 {
    ans += a
    new_a = (-p)%a
    if new_a < 0 {
      new_a += a
    }
    new_p = a
    a = new_a
    p = new_p
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
