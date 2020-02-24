/ pe008

input: raze read0 `:../input.txt
max {prd"J"$/:13#x _ y}\:[til -13 + count input; input]
max prd each "J"$/:/:input (til -13+count input) +\: til 13
