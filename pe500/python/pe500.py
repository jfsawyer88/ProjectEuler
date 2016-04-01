#!/usr/bin/python

# project euler
# problem 500!!!

import bisect
import gmpy2
import math
import time

def pe500(e, M):
    prd = 1
    np = 3
    queue = [np]
    while e-1 > math.log(math.log(np, 2), 2):
        if queue[0] == np:
            bisect.insort(queue, np**2)
            np = gmpy2.next_prime(np)
            bisect.insort(queue, np)
        else:
            bisect.insort(queue, queue[0]**2)
        prd = (prd * queue.pop(0)) % M
        e -= 1
    return (pow(2, 2**e - 1, M) * prd) % M


t0 = time.time()
output = pe500(500500, 500500507)
print "timing:", time.time() - t0
print "output:", output
