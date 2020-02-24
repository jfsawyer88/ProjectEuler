// euler 006

sq:{x*x}

pe006_1: {(sq sum 1+til x) - sum sq 1+til x}
pe006_2: {(sq (x*(1+x)) div 2) - (x*(1+x)*(1+2*x)) div 6}

\t:1000 pe006_1 1000
\t:1000 pe006_2 1000
