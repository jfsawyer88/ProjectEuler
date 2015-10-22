#!/usr/bin/python

## Project Euler
## Smallest Multiple

import time

def gcd(a, b):
    while a != 0:
        c = a
        a = b % a
        b = c
    return b

def lcm(a, b):
    return (a * b) / gcd(a, b)

t0 = time.time()
N = 20

ans = 1
for n in xrange(1, N + 1):
    ans = lcm(ans, n)

print "Execution time:", time.time() - t0
print "Output:", ans
