#!/home/ubuntu/j/bin/jconsole

NB. Project Euler
NB. Self Powers

time=: ": 6!:2 'output=: ":(1e10|+)/ (1e10&|@^)~"0 >: i. 1000'

stdout 'Execution time: ', time, LF
stdout 'Output: ', output, LF

exit''
