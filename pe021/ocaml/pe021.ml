(** Project Euler **)
(** Amicable Numbers **)


let t0 = Unix.gettimeofday ()

(** input **)
let limit = 10000;;

(** sum of divisors in array **)
let sod_array = Array.make limit 0;;
for i = 1 to limit - 1 do
  begin
    let j = ref (2 * i) in
    while !j < limit do
      Array.set sod_array !j (i + Array.get sod_array !j);
      j := i + !j;
    done
  end
done
;;

(** sum of divisors function **)
(** all values for n < 10000 stored **)
let sum_of_divisors n =
  if n < limit then
    Array.get sod_array n
  else
    let rec aux d c =
      if d * d > n then c
      else
	if 0 = n mod d then
	  if d * d = n then
	    aux (succ d) (c + d)
	  else
	    aux (succ d) (c + d + (n / d))
	else
	  aux (succ d) c
    in
    aux 2 1
;;


let res = ref 0;;
for n = 1 to limit - 1 do
  let m = sum_of_divisors n in
  if (n = sum_of_divisors m) && (n <> m) then
    res := !res + n;
done
;;


Printf.printf "Execution time: %f seconds\n" (Unix.gettimeofday () -. t0);;
Printf.printf "Output: %d\n" !res
