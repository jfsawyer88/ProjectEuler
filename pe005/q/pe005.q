/gcd:{[a;b] while[a<>b; $[a>b;a-:b;b-:a]]; a}
gcd:{[a;b]
    while[0<>b;
        t:b;
        b:a mod b;
	a:t
    ];
    a
 }
lcm:{[a;b] (a*b) div gcd[a;b]}

gcd_rec:{[a;b] $[0=b; a; gcd_rec[b; a mod b]]}
lcm_rec:{[a;b] (a*b) div gcd_rec[a;b]}

t0:.z.N
(lcm/) 1+til 20
.z.N-t0

t0:.z.N
(lcm_rec/) 1+til 20
.z.N-t0
\\
