NB.#!/home/ubuntu/j/bin/jconsole

NB. Project Euler
NB. Summation of Primes

time=: ": 6!:2 'output=: ": +/ i.&.(p:inv) 2e6'

stdout 'Execution time: ', time, LF
stdout 'Output: ', output, LF

NB.exit''
