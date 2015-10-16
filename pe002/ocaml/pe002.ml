(** Project euler **)
(** Even Fibonacci Numbers **)

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



let pe002 lim =
  let rec loop f1 f2 res =
    if lim < f2 then res
    else
      if 0 = f2 mod 2 then
	loop f2 (f1 + f2) (res + f2)
      else
	loop f2 (f1 + f2) res
  in
  loop 1 2 0
;;

time (fun () -> (pe002 4000000)) 1
