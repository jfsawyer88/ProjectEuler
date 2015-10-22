#!/usr/bin/python

## Project Euler
## Self Powers

import time

t0 = time.time()

ans = sum(map((lambda x: pow(x, x, 10000000000)), xrange(1,1001))) % 10000000000

print "Execution time:", time.time() - t0
print "Output:", ans
