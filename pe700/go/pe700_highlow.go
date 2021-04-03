package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()
  a := int64(1504170715041707)
  p := int64(4503599627370517)
  lo := a
  hi := a
  ans := a

  var next int64
  for lo > 0  {
    next = (lo+hi)%p
    if next < lo {
      lo = next
      ans += lo
    } else {
      hi = next
    }
  }

  fmt.Printf("Result is %v\n", ans)
  fmt.Printf("Timing: %v", time.Since(start))
}
