################################
## Project Euler 512          ##
## Sums of totients of powers ##
## +/5 p: >:+:i.-: 5e8        ##
################################

import time

t0 = time.time()

N = 10 ** 6

phi = range(N + 1)
i = 2
while i < N + 1:
    if i == phi[i]: # if i is prime
        phi[i] -= 1
        j = 2 * i
        while j < N + 1:
            phi[j] = (phi[j] // i) * (i - 1)
            j += i
    i += 1

print time.time() - t0, 'seconds'
print sum(phi[1::2])
