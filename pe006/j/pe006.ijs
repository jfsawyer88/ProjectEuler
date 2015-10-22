#!/home/ubuntu/j/bin/jconsole

NB. Project Euler
NB. Sum Square Difference

time=: ": 6!:2 'output=: ": (([: *: +/) - ([: +/ *:)) >:i.100'

stdout 'Execution time: ', time, LF
stdout 'Output: ', output, LF

exit''
