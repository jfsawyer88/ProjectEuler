(** Project Euler **)
(** Self Powers **)

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


let mpow a b n =
  let rec aux a b n res =
    if b = 0 then res
    else aux a (b - 1) n ((a * res) mod n)
  in
  aux a b n 1
;;

let rec  mpow a b n =
  match b with
  | 0 -> 1
  | 1 -> a mod n
  | b -> (mpow a (b mod 2) n) * (mpow ((a * a) mod n) (b / 2) n)
;;


let lmsum l n =
  let rec aux l res =
    match l with
    |    []    -> res
    | hd :: tl -> aux tl ((hd + res) mod n)
  in
  aux l 0
;;

let range a b =
  let rec aux a b res =
    if b < a then res
    else aux a (b - 1) (b :: res)
  in
  aux a b []
;;


let pe048 lim =
  lmsum (List.rev_map (fun n -> (mpow n n 10000000000)) (range 1 lim)) 10000000000
;;


time (fun () -> (pe048 1000)) 1
