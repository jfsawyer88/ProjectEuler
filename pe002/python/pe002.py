#!/usr/bin/python

## Project Euler
## Even Fibonacci Number

import time

t0 = time.time()

ans = 0
f1 = 1
f2 = 2
while f2 < 4000000:
    if 0 == f2 % 2:
        ans += f2
    f2 = f1 + f2
    f1 = f2 - f1

print "Execution time:", time.time() - t0
print "Output:", ans
