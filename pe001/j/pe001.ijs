#!/home/ubuntu/j/bin/jconsole

time=: ": 6!:2 'out=: ":+/ (#~ 0 = [:*/ 3 5 |"0 1 ]) i.1000'

stdout 'Execution time: ', time, LF
stdout 'Output: ', out, LF

exit ''
