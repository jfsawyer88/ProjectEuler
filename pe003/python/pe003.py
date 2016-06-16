#!/usr/bin/python

## Project Euler
## Largest Prime Factor

import time
t0 = time.time()

N = 600851475143
factor = 2
last_factor = 1
#while N > 1:
while factor*factor < N:
    if 0 == N % factor:
        last_factor = factor
        N /= factor
        while 0 == N % factor:
            N /= factor
    factor += 1

print "Execution time:", time.time() - t0
print "Output:", last_factor
