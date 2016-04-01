#!/usr/bin/python

## Project Euler
## Truncatable Primes

import time
t0 = time.time()

def is_prime(n):
    if 0 == n % 2:
        return False
    i = 3
    while i < 1 + n ** 0.5:
        if 0 == n % i:
            return False
        i += 2
    return True

ans = 0
good_digits = {1, 2, 3, 5, 7, 9}
prev_ltp = {2, 3, 5, 7}
prev_rtp = {2, 3, 5, 7}
n_digits = 1

while prev_ltp and prev_rtp:

    ltp = set()
    for a in good_digits:
        for b in prev_ltp:
            n = (a * 10 ** n_digits) + b
            if is_prime(n):
                ltp.add(n)

    rtp = set()
    for a in prev_rtp:
        for b in good_digits:
            n = (a * 10) + b
            if is_prime(n):
                rtp.add(n)

    ans += sum(ltp & rtp)
    prev_ltp = ltp
    prev_rtp = rtp
    n_digits += 1

print "Execution time:", time.time() - t0
print "Output:", ans
