collatz_array:1000#0
collatz_array[1]:1

collatz:{$[x mod 2; 1+3*x; x div 2]}

collatz_chain: collatz\[1<>;]
