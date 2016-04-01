#!/usr/bin/python

## Project Euler
## Champernowne's Constant

import time
t0 = time.time()

def digits(n, b=10):
    res = []
    while 0 < n:
        d = n % b
        n = (n - d)/b
        res.append(d)
    res.reverse()
    return res

temp = [9 * 10**i * (i + 1) for i in xrange(1, 18)]
intervals = [1, 10]
for i in temp:
    intervals.append(intervals[-1] + i)

to_find = [10**i for i in xrange(1, 18)]
ans = 1
for n in to_find:
    for i in xrange(len(intervals)):
        if intervals[i] > n:
            i -= 1
            break
    q = (n - intervals[i]) / (i + 1)
    r = (n - intervals[i]) % (i + 1)
    # answers original problem
    #ans *= digits(q + 10**i)[r]
    #print digits(q + 10**i)[r], q+10**i, r
    # gives result of concatenating the digits instead
    # of multiplying them (first 18, f brute force)
    # under 100 micro-seconds
    ans = ans * 10 + digits(q + 10**i)[r]

print "Execution time:", time.time() - t0
print "Output:", ans
