#!/home/ubuntu/j/bin/jconsole

NB. Project Euler
NB. Smallest Multiple

time=: ": 6!:2 'output=: ": *./ >: i. 20'

stdout 'Executiont time: ', time, LF
stdout 'Output: ', output, LF

exit''
