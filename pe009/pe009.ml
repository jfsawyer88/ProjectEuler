(** Project Euler **)
(** Special Pythagorean Triplet **)

(** timing function **)
(*#load "unix.cma";;*)
let time f number =
  let t = Unix.gettimeofday () in  (** get the current time **)
  let time_unit = match number with
    | 1 -> ""
    | 1000 -> "milli-"
    | 1000000 -> "micro-"
    | 1000000000 -> "nano-"
    | _ -> "1/" ^ (string_of_int number) ^ "-"
  in
  let rec loop i =
    if i = number then f ()        (** evaluate f number times **)
    else
      begin
	ignore (f ());
	loop (i + 1)
      end
  in
  let res = loop 1 in              (** keep the result to print **)
  Printf.printf "Execution time: %f %sseconds\n"
                (Unix.gettimeofday () -. t)
		time_unit;
  Printf.printf "Output: %d\n" res;
;;


(** returns the integer square root of n **)
(** i.e. floor(sqrt(n)) **)
let  isqrt n =
  let rec loop xk = 
    let xk1 = (xk + (n / xk)) / 2 in  (** computs the next x value **)
    if 1 >= (xk - xk1) then xk1       (** compares it with our current x value **)
    else loop xk1
  in
  loop n
;;


let pe009 () =
  let rec loop a b =
    let c = isqrt ((a * a) + (b *b)) in
    if b > a then loop (a + 1) 1
    else
      if ((a * a) + (b * b)) <> (c * c) then loop a (b + 1)
      else
	if (a + b + c) = 1000 then a * b * c
	else loop a (b + 1)
  in
  loop 1 1
;;

time pe009 1000
