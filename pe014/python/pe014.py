#!/usr/bin/python

## Project Euler
## Longest Collatz Sequence

import time

t0 = time.time()
N = 1000000

def collatz(n):
    if 0 == n % 2:
        return n / 2
    else:
        return 3 * n + 1


collatz_array = [0] * N
collatz_array[1] = 1
m = 1
for i in xrange(2, N):
    s = collatz(i)
    c = 1
    while s >= i:
        s = collatz(s)
        c += 1
    collatz_array[i] = c + collatz_array[s]
    if collatz_array[m] < collatz_array[i]:
        m = i


print "Execution time", time.time() - t0
print "Output:", m
