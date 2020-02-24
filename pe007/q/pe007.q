// pe007

is_prime:{[n]
 if[n<2; :0b];
 if[n<4; :1b];
 if[0=n mod 2; :0b];
 if[n<9; :1b];
 if[0=n mod 3; :0b];
 r: floor sqrt 51;
 f:5;
 while[f <= r;
  if[0 = n mod f  ; :0b];
  if[0 = n mod f+2; :0b];
  f+:6
  ];
 1b
 }

pe007:{[goal]
 n:1;
 nprimes:1;
 while[nprimes<goal;
  n+:2;
  if[is_prime[n]; nprimes+:1]
  ];
 n
 }

\t show pe007 10001






