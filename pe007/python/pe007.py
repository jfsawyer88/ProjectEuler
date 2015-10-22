#!/usr/bin/python

## Project Euler
## 10001st Prime

import time

t0 = time.time()


def isqrt(n):
    xk = n
    xk1 = (xk + (n / xk)) / 2
    while xk1 < xk:
        xk = xk1
        xk1 = (xk + (n / xk)) / 2
    return xk

def prime_pi(n):
    "count primes below n"
    r = isqrt(n)
    V = [n//i for i in range(1,r+1)]
    V += list(range(V[-1]-1,0,-1))
    S = {i:i-1 for i in V}
    for p in range(2,r+1):
        if S[p] > S[p-1]:  # p is prime
            sp = S[p-1]  # sum of primes smaller than p
            p2 = p*p
            for v in V:
                if v < p2: break
                S[v] -= (S[v//p] - sp)
    return S[n]

def primes_upto(limit):
    "returns all primes up to limit"
    is_prime = [False] * 2 + [True] * (limit - 1) 
    for n in range(int(limit**0.5 + 1.5)): # stop at ``sqrt(limit)``
        if is_prime[n]:
            for i in range(n*n, limit+1, n):
                is_prime[i] = False
    return [i for i, prime in enumerate(is_prime) if prime]

def primes_between(a, b):
    ps = primes_upto(b)
    P = [1] * (b - a)
    for p in ps:
        for k in xrange((-a)%p, b-a, p):
            if k + a <= p: continue
            P[k] = 0
    return [a+i for i, prime in enumerate(P) if prime]


N = 10001

i = 1
pp = 0
while pp < N:
    i *= 2
    prev_pp = pp
    pp = prime_pi(i)

ans = primes_between(i/2, i)[N-prev_pp-1]

print "Execution time:", time.time() - t0
print "Output:", ans
