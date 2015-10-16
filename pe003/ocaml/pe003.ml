(** Project Euler **)
(** Largest Prime Factor **)


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


let pe003 n =
  let rec loop i n =
    if i = n then n
    else
      loop (i + 1) (ddiv n i)
  in
  loop 2 n
;;


time (fun () -> (pe003 600851475143)) 1
