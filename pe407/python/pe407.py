#!/usr/bin/python
# pe 407


import time
t0 = time.time()

N = 1000000

# number of factors of p in n
def nf(n, p):
    cnt = 0
    while 0 == n % p:
        n /= p
        cnt += 1
    return cnt

# largest prime factor
LPF = [0] + [1] * N
for p in xrange(2, N + 1):
    # if p is prime
    if 1 == LPF[p]:
        LPF[p] = p
        for n in xrange(2 * p, N + 1, p):
            pk = p ** nf(n, p)
            if LPF[n] < pk:
                LPF[n] = pk

output = 0
for n in xrange(2, N + 1):
    if n == LPF[n]:
        output += 1
    else:
        for a in xrange(n - LPF[n], n/2 - 1, -LPF[n]):
            if (a + 1) == ((a + 1) ** 2) % n:
                output += (a + 1)
                break
            elif a == (a ** 2) % n:
                output += a
                break

print 
print "Output:", output
print "Timing:", time.time() - t0
