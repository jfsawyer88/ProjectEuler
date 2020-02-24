// project euler 021
// amicable numbers

sod_sub: {indx:range[2*x;count y;x]; (1+x;@[y;indx; +; x])}.
sod: {last (x div 2) sod_sub/ (1; x#0)}
pe021: {[N] sum where {y and not x}. {(first x) =/: 1_x} 2 (sod N)\ til N}
\t pe021 10000

//////////////////////
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

divisors: {{asc (raze/) (*/:/) prds each 1,/: value x #' key x} count each group factorize x}
spd: {$[2>x;0;sum -1_divisors x]} each
amicableq: {{01b~(first x) = 1_x} 2 spd\ x} each
\t sum where amicableq til 10000
