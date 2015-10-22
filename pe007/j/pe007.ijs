#!/home/ubuntu/j/bin/jconsole

NB. Project Euler
NB. 10001st Prime

time=: ": 6!:2 'output=: ": p:<:10001'

stdout 'Execution time: ', time, LF
stdout 'Output: ', output, LF

exit''
