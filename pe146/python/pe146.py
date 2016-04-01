#!/usr/bin/python
# pe 146

import time
import gmpy2

t0 = time.time()

limit = 150 * 10**6
diff = [2, 6, 8, 12, 26]


output = 0
for n in xrange(0, limit, 10):
    p = n**2 + 1
    if gmpy2.is_prime(p):
        next_ps = [p + d for d in diff]
        goodn = True
        for np in next_ps:
            p = gmpy2.next_prime(p)
            if p != np:
                goodn = False
                break
        if goodn:
            output += n



print output
print time.time() - t0
