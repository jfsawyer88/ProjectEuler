#!/usr/bin/python
# pe521

import gmpy2
import time

t0 = time.time()
N = 1000000000000
M = 1000000000

def pe521(n):
    r = gmpy2.isqrt(n)
    V = [n//i for i in range(1,r+1)]
    V += list(range(V[-1]-1,0,-1))
    np =   {i:i-1          for i in V}
    sp =   {i:i*(i+1)//2-1 for i in V}
    ssmp = {i:i*(i+1)//2-1 for i in V}
    for p in range(2,r+1):
        if np[p] > np[p-1]:  # p is prime
            p2 = p*p
            for v in V:
                if v < p2: break
                np[v] -=   (np[v//p] - np[p-1])
                sp[v] -= p*(sp[v//p] - sp[p-1])
                ssmp[v] -= p * ( (sp[v//p] - np[v//p]) - (sp[p-1] - np[p-1]) )
    return ssmp[n]

output = pe521(N) % M

print "Output:", output
print "timing:", time.time() - t0
