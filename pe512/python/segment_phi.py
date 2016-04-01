#!/usr/bin/python

################################
## phi of all below N         ##
## variant of segmented sieve ##
################################

import time
t0 = time.time()

def primes_below(limit):
    '''gives a list of all primes below the given limit'''
    is_prime = [False] * 2 + [True] * (limit - 2)
    for n in range(int(1 + limit**0.5)):
        if is_prime[n]:
            for i in range(n*n, limit, n):
                is_prime[i] = False
    return [i for i, prime in enumerate(is_prime) if prime]

def full_div(n, p):
    '''remove all factors of p from n'''
    while 0 == n % p:
        n /= p
    return n

#N = 1 + (5 * 10 ** 8)
N = 1 + 10 ** 6
# get an initial count for values below sqrt(N)
M = int(N ** 0.5)
phi = range(M)
i = 2
while i < M:
    if i == phi[i]: # if i is prime
        phi[i] -= 1
        j = 2 * i
        while j < M:
            phi[j] = (phi[j] // i) * (i - 1)
            j += i
    i += 1

output = sum(phi[1::2])

segment_size = 2**15 # 32KB segments
primes = primes_below(int(N ** 0.5))
low = int(N ** 0.5)
next_indx = [(-low) % p for p in primes]

k = 0 # to keep track of how long we have left in console
while low < N:
    if 0 == k % 100: print 100 * low / float(N+1), '% done at', time.time() - t0, 'seconds.'
    phi = [low + n for n in range(segment_size)] # initialize current segment
    T = [low + n for n in range(segment_size)] # to divide fully by primes
    i = 0
    while i < len(primes):
        while next_indx[i] < segment_size:
            phi[next_indx[i]] = (phi[next_indx[i]] // primes[i]) * (primes[i] - 1)
            T[next_indx[i]] = full_div(T[next_indx[i]], primes[i])
            next_indx[i] += primes[i]
        i += 1
    i = 0
    t = low
    while (i < segment_size) and (t < N):
        if T[i] != 1: # if there is a factor of t > sqrt(N) the fix phi
            phi[i] = (phi[i] // T[i]) * (T[i] - 1)
        if 1 == t % 2:
            output += phi[i]
        i += 1
        t += 1
    next_indx = [indx - segment_size for indx in next_indx]
    low += segment_size
    k += 1

print time.time() - t0
print output
