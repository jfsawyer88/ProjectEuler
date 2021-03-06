(** Project Euler **)
(** Distinct Primes Factors **)

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

(** divide out all factors of n from a **)
let rec ddiv a n =
  if 0 <> a mod n then a
  else
    ddiv (a / n) n
;;


(** count distinct factors **)
let cdf m =
  let rec loop n p c =
    if n = 1 then c
    else
      if p * p > n then succ c (** breaking on this speeds things waay up **)
      else
	if 0 = n mod p then loop (ddiv (n / p) p) (succ p) (succ c)
	else loop n (succ p) c
  in
  loop m 2 0
;;


let pe047 () =
  let rec loop n c =
    (*Printf.printf "n:%d, c:%d\n" n c;*)
    if c = 4 then n - 4
    else
      if 4 = cdf n then loop (succ n) (succ c)
      else loop (succ n) 0
  in
  loop 2 0
;;

time pe047 1
