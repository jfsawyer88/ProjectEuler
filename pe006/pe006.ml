(** Project Euler **)
(** Sum Square difference **)


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
  let res = loop 0 in              (** keep the result to print **)
  Printf.printf "Execution time: %f %sseconds\n"
                (Unix.gettimeofday () -. t)
		time_unit;
  Printf.printf "Output: %d\n" res;
;;


(** sum of squares **)
let sumsqrs n =
  let rec loop i res =
    if i > n then res
    else
      loop (i + 1) (res + i * i)
  in
  loop 1 0
;;

(** square of sum **)
let sqrsum n =
  let rec loop i res =
    if i > n then res
    else
      loop (i + 1) (res + i)
  in
  let out = loop 1 0 in
  out * out
;;


let pe006_1 n =
  (sqrsum n) - (sumsqrs n)
;;

let pe006_2 n =
  let sum = (n * (n + 1)) / 2 in
  (sum * sum) - ((2 * n + 1) * (n + 1) * n) / 6
;;


time (fun () -> (pe006_1 100)) 1000000;;
time (fun () -> (pe006_2 100)) 1000000000
