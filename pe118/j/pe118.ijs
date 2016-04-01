#!/home/ubuntu/j/bin/jconsole

np=: ([:+/ 1 p: 10 #. i.@:!@:# A. ])
draw=: [: ([:<0:)`([:(<"1) 0,. [:#: [: i. 2 ^ <:)@.(1~:]) #
foo=: [:|:@:> draw (#;(-.@:[ # ])) L: 0 ]
np2=: [:; [:np each {:
nps=: (((1:)`(([:(([:; [:$: each {.) +/ . * np2) foo))@.(0 ~: #)))

time=: ": 6!:2 'out=: ": nps 1 2 3 4 5 6 7 8 9'

stdout 'Execution time: ', time, LF
stdout 'Output: ', out, LF

exit''
