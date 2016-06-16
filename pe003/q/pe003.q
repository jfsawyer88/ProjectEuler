t0:.z.N

N:600851475143
pe003:{[N]

    // currenct factor is initialized at 2
    factor:2;

    // largest factor is initialized at 1
    last_factor:1;

    // loop it up!
    while[ N > 1;
        if[ 0 = N mod factor;
	    last_factor:factor;
	    N:N div factor;
	    while[0 = N mod factor;
	        N:N div factor
		]
	];
	factor+:1
	];

    last_factor
 }


pe003 N
.z.N-t0
\\
