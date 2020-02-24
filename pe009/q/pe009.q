factorize:{[N]
 / first test factor and list of all factors
 f:2; fs:`long$();
 / end when all factors are found or
 / we reach the sqrt of current N
 while[f < 1+sqrt N;
  while[0 = N mod f;
   fs,:f;      / append factor
   N: N div f  / reduce N
   ];
  f+:1
  ];
 fs,$[1=N;();N]
 }


ndiv: {prd 1+count each group factorize x}
divisors: {{asc (raze/) (*/:/) prds each 1,/: value x #' key x} count each group factorize x}
gcd: {first {0<>last x} ({y, x mod y}.)/ x}


pe009:{[s]
 if[s mod 2; :()];     / s must be even
 ns: divisors s div 2; / n divides (s/2)
 ns: 1_ns where ns < ceiling sqrt s%2; / n < sqrt(s/2) and n<>1
 nks: raze {[s;n]      / k = (n+m)
  /k_max: s div 2 * n;  / k is factor of (s/(2*n))
  /code below guarantees the conidtion above
  k_max: ({$[x mod 2; x; x div 2]}/) s; / k must be odd - gcd(n,m)=1
  k_max: first ({((first x) div gcd x), last x}/) k_max, n; / gcd(n,k)=1
  ks:divisors k_max;   / possible k's are divisors of k_max
  n ,/: ks where ks within {x, 2*x} n  / we must have n < k < 2n
  }[s] each ns;
 nmd: ({x,(y-x), z div 2*x*y}.) each nks ,\: s; / with n and k compute n, m, and d
 ({z * ((x*x)-(y*y)) , (2*x*y) , (x*x)+(y*y)}.) each nmd / now compute a, b, and c
 }


\t:1000 prd first pe009 1000

pe009 38294381024
pe009 40
pe009 4050
pe009 38294381026


//////////////////////////////////////
factorize_imp:{[N]
 / first test factor and list of all factors
 f:2; fs:`long$();
 / end when all factors are found or
 / we reach the sqrt of current N
 while[f < 1+sqrt N;
  while[0 = N mod f;
   fs,:f;      / append factor
   N: N div f  / reduce N
   ];
  f+:1
  ];
 fs,$[1=N;();N]
 }


div_out: {[N;f;cnt] (N div f; f; 1+cnt)}.
cond1: {[N;f;cnt] 0 = N mod f}.
fact_sub1:{[N;f;cnt] {(x;z#y)}. cond1 div_out/ (N;f;cnt)}
fact_sub2:{[N;fs;f] {(x[0];y,x[1])}[fact_sub1[N;f;0];fs],1+f}.
cond2: {[N;fs;f] f < 1+sqrt N}.
factorize: {[N] {x[1],(1<>x[0])#x[0]} cond2 fact_sub2/ (N;`long$();2)}

\t factorize 382943810279
\t factorize_imp 382943810279

\t factorize 38294381026
\t factorize_imp 38294381026

\t factorize 100000000003
\t factorize_imp 100000000003