#!/usr/bin/python

## Project Euler
## Multiples of 3 and 5

import time

t0 = time.time()

ans = 0
for i in xrange(1000):
    if (0 == i % 3) or (0 == i % 5):
        ans += i

print "Execution time:", time.time() - t0
print "Output:", ans
