primepi:{first "J"$system"primecount.exe ",string x};

tosieve:{[N;p]p*p _ til ceiling N%p};
sieve_p:{[N;v;p] .debug.input:(v;p);@[v;tosieve[N;p];:;0b]};
sieve:{[N;v;p]$[v[p];sieve_p[N;v;p];v]};

primes_below:{[N] v:@[N#1b;0 1;:;0b];where v sieve[N]/ til ceiling sqrt N};

// number of semi-primes below N:
\ts 0N!sum {x+neg til count x} primepi each N div primes_below ceiling sqrt N:100000000
