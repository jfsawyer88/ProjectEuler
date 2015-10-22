#!/usr/bin/python

## Project Euler
## Sum Square Difference

import time

t0 = time.time()

out = sum(xrange(1, 101)) ** 2 - sum(map(lambda x: x**2, xrange(1, 101)))

print "Execution time:", time.time() - t0
print "Output:", out
