#!/home/ubuntu/j/bin/jconsole

NB. project euler
NB. problem 525

9!:11 ]12

sin=: 1&o.
cos=: 2&o.

a=: 2
b=: 4

 r=: ((a*b) % [: +/&.:*: (a, b) * sin, cos)"0
dr=: (((a*b*(b+a)*(b-a) * sin * cos)) % 3^~[: +/&.:*: (a, b) * sin, cos)"0

f=: ([: +/&.:*: r, dr)"0
L=: ((% * [: +/ [: f [*(%~i.)@:])&1000)"0

dist=: [:(+/&.:*:)-
+/2 dist/\ (r,.L) 2p1*(%~i.@:>:) 100