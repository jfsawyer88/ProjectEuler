######################################
## Segmented Sieve of Eratosthenes  ##
## used here to give primes below N ##
######################################

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


N = 10 ** 8 # input size
segment_size = 2**15 # segment size, 32KB

primes = primes_below(int(N ** 0.5))
output = [p for p in primes]
low = int(N ** 0.5) # starting point for segmentation
next_indx = [(-low) % p for p in primes] # indicies of next primes in upcoming sieve

while low < N:
    sieve = [True] * segment_size # initialize segment
    i = 0
    while i < len(primes):
        while next_indx[i] < segment_size:
            sieve[next_indx[i]] = False
            next_indx[i] += primes[i]
        i += 1
    i = 0
    t = low
    while (i < segment_size) and (t < N):
        if sieve[i]:
            output.append(t)
        i += 1
        t += 1
    next_indx = [indx - segment_size for indx in next_indx]
    low += segment_size


print time.time() - t0
print len(output)
