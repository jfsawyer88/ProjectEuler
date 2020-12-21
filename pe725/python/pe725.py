#!/usr/bin/python

## Project Euler Problem 725
## Digit sum numbers

import time
t0 = time.time()

def partitions(n):
    ## returns the partitions of n
    ans = []
    for k in range(1, n+1):
        ans += (partitions_sub(n, k))
    return ans

def partitions_sub(n, k):
    ## returns the partitions of n with largest summand=k
    if k==1: return [[1]*n]
    if k==n: return [[n]*1]
    ans = []
    for i in range(1, 1+min(k, n-k)):
        for p in partitions_sub(n-k, i):
            ans.append(p + [k])
    return ans

def valid_partitions(n, k):
    ## partitions of k with count of summands < n, with k
    res = []
    for p in partitions(k):
        if len(p) < n:
            p.append(k)
            res += [[0]*(n-len(p)) + p]
    return res

def factorial(n):
    ## returns the factorial of n in factorized form
    ## represented as a dictionary of primes -> exponents
    ans = {}
    for p in primes_below(n+1):
        ans[p] = factorial_multiplicity(n, p)
    return ans

def factorial_multiplicity(n, p):
    ## return number of time p divides n! (n factorial)
    e = 0
    pp = p
    while n//pp > 0:
        e += n//pp
        pp *= p
    return e

def primes_below(N):
    ## returns list of primes < N
    is_prime = [True] * N
    is_prime[0] = False
    is_prime[1] = False
    p = 2
    while p*p < N:
        if is_prime[p]:
            for n in range(2*p, N, p):
                is_prime[n] = False
        p+=1
    primes = []
    for p in range(N):
        if is_prime[p]:
            primes.append(p)
    return primes

def overcounting_factor(vp):
    ## returns the overcounting factor of valid partition vp
    map = {}
    for d in vp:
        if d in map:
            map[d] += 1
        else:
            map[d] = 1
    ans = {}
    for d in map:
        ans = multiply(ans, factorial(map[d]))
    return ans

def multiply(a, b):
    ## multiply numbers a by b
    ## a and b are fully factorized, represented as dictionaries of primes -> exponents
    res = dict(a)
    for p in b:
        if p in a:
            res[p] += b[p]
        else:
            res[p] = b[p]
    return res

def divide(a, b):
    ## multiply numbers a by b
    ## a and b are fully factorized, represented as dictionaries of primes -> exponents
    ## assumes b | a
    res = dict(a)
    for p in b:
        res[p] -= b[p]
    return res

def factorize(n):
    ans = {}
    p = 2
    while p*p <= n:
        if 0 == n%p:
            ans[p] = 1
            n//=p
            while 0 == n%p:
                ans[p] += 1
                n//=p
        p+=1
    # if n is not reduced to 1 the remaining value is prime itself
    if 1 < n:
        ans[n] = 1
    return ans

def evaluate(a, q):
    ## convert dictionary->exponent dictionary to number mod q
    res = 1
    for p in a:
        res *= pow(p, a[p], q)
    return res

def S_sub(n, k, q):
    sum=0
    full = multiply(factorize(2*k), factorial(n-1))
    for p in valid_partitions(n, k):
        sum = (sum + evaluate(divide(full, overcounting_factor(p)), q)) % q
    return sum

def S(n, q):
    res=0
    for k in range(1, 10):
        res+=S_sub(n, k, q)
    ones=0
    for i in range(n):
        ones = (ones + pow(10, i, q)) % q
    return (res * ones) % q

ans = S(2020, 10**16)
print("Execution time:", time.time() - t0)
print("Output:", ans)
