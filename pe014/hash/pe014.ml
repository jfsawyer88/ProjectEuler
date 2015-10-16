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


let pe014 n =
  let l = Hashtbl.create n in
  Hashtbl.add l 1 1;
  let rec aux1 s m c =
    if s < m then c + Hashtbl.find l s
    else aux1 (collatz s) m (succ c)
  in
  let rec aux2 i out =
    if i > n then out
    else
      begin
	Hashtbl.add l i (aux1 i i 0);
	if (Hashtbl.find l out) < (Hashtbl.find l i) then aux2 (succ i) i
	else aux2 (succ i) out
      end
  in
  aux2 2 1
;;

time (fun () -> (pe014 1000000)) 1
