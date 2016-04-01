#!/home/ubuntu/j/bin/jconsole

n=: 10^12x NB. in ~1.3ms
rot=: 4 : '((2 2$0 _1 1 0x)&(+/ .*))&.(-&x)  y'
lpb=: [:x:2x^[:<.2x^.]
path=: <.2^.(-~2*lpb)^:a: n
tab=: (rot&0 0)^:(<>:{.path) 0 1
stdout 'Execution time: ', (":6!:1''), LF
stdout 'Output: ' , (' ,' stringreplace ":rot/path{tab), LF

exit''
