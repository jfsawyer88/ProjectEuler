#!/home/ubuntu/j/bin/jconsole

time=: ": 6!:2 'out=: ":+/ (#~0 = 2|]) (,[:+/_2{.])^:(4000000 > {:)^:_] 1 1'

stdout 'Execution time: ', time, LF
stdout 'Output: ', out, LF

exit ''
