(** Project Euler **)
(** Smallest Multiple **)


(** timing function **)
(*#load "unix.cma";;*)
let time f number =
  let t = Unix.gettimeofday () in  (** get the current time **)
  let time_unit = match number with
    | 1 -> ""
    | 1000 -> "milli-"
    | 1000000 -> "micro-"
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
  let res = loop 0 in              (** keep the result to print **)
  Printf.printf "Execution time: %f %sseconds\n"
                (Unix.gettimeofday () -. t)
		time_unit;
  Printf.printf "Output: %d\n" res;
;;


(** greatest common divisor **)
let rec gcd a b =
  if b = 0 then a
  else gcd b (a mod b)
;;

(** least common multiple **)
let lcm a b =
  (a * b) / (gcd a b)
;;

(** smallest multiple of all i < n **)
let pe005 n =
  let rec loop i res =
    if i > n then res
    else
      loop (i + 1) (lcm i res)
  in
  loop 1 1
;;

time (fun () -> (pe005 20)) 1000000
