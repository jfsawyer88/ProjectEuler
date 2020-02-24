// pe016

condition: {0 <> count first x}
carry_sub: {((-1_x);(((last x)+first y div 10),(first y) mod 10),1_y)}.
drop_leading_zero:{$[0=first x;1_x;x]}
digits: {"J"$/:string x}
expand_first_digit: {(digits first x), 1_x}
fmt_input: {(-1_x;enlist last x)}
carry: {expand_first_digit last condition carry_sub/ fmt_input  x} 

\t sum 1000 {carry 2*x}/ enlist 1
