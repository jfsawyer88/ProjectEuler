t0:.z.N
sum {x where 0 = x mod 2} -1_{$[4000000>last x;x,sum -2#x;x]}/[1 2]
.z.N-t0

t0:.z.N
fibs: -1_{$[4000000>last x;x,sum -2#x;x]}/[1 2]
sum fibs where 0 = fibs mod 2
.z.N-t0

t0:.z.N
fib:1 2
while[ 4000000>last fib;
 fib,:sum -2#fib
 ]
fib:-1_fib
sum fib where 0 = fib mod 2
.z.N-t0

t0:.z.N
sum {x where 0 = x mod 2} -1_{x,sum -2#x}/[{4000000>last x};1 2]
.z.N-t0

t0:.z.N
fibs:-1_{x,sum -2#x}/[{4000000>last x};1 2]
sum fibs where 0 = fibs mod 2
.z.N-t0

t0:.z.N
fib:{x,sum -2#x}/[;1 2]
fibs:fib{4000000>last x}
sum fibs where not fibs mod 2
.z.N-t0
\\
