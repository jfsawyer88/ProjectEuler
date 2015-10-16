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


let collatzl n l =
  let rec aux s c =
    if s < n then c + List.nth l (n - s - 1)
    else aux (collatz s) (succ c)
  in
  aux n 0
;;


let pe014 n =
  let rec aux l m i =
    if i > n then m
    else
      match collatzl i l with
      | li when li > List.nth l (i - m - 1) -> aux (li :: l) i (succ i)
      | li                                  -> aux (li :: l) m (succ i)
  in
  aux [1] 1 2
;;


time (fun () -> (pe014 1000000)) 1;;
