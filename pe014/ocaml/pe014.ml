(** Project Euler **)
(** Longest Collatz Sequence **)

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

let collatz n =
  if 0 = n mod 2 then n / 2
  else 3 * n + 1
;;

let collatz_plux a n =
  let rec aux s c =
    if s < n then c + Array.get a s
    else aux (collatz s) (succ c)
  in
  aux n 0
;;


let pe014 n =
  let collatz_array = Array.make n 0 in
  Array.set collatz_array 1 1;
  let max_chain = ref 1 in
  for i = 2 to n - 1 do
    Array.set collatz_array i (collatz_plux collatz_array i);
    if (Array.get collatz_array !max_chain) < (Array.get collatz_array i) then
      max_chain := i
  done;
  !max_chain
;;


time (fun () -> (pe014 1000000)) 1;;
