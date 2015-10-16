(** Project Euler **)
(** Summation of Primes **)

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



(* first define some iterators *)
let fold_iter f init a b =
  let rec aux acc i =
    if i > b
    then (acc)
    else aux (f acc i) (succ i)
  in
  aux init a
;;
 
let fold_step f init a b step =
  let rec aux acc i =
    if i > b
    then (acc)
    else aux (f acc i) (i + step)
  in
  aux init a ;;

 
(* remove a given value from a list *)
let remove li v =
  let rec aux acc = function
    | hd::tl when hd = v -> (List.rev_append acc tl)
    | hd::tl -> aux (hd::acc) tl
    | [] -> li
  in
  aux [] li ;;

 
(* the main function *)
let primes n =
  let li =
    (* create a list [from 2; ... until n] *)
    List.rev(fold_iter (fun acc i -> (i::acc)) [] 2 n)
  in
  let limit = truncate(sqrt(float n)) in
  fold_iter (fun li i ->
              if List.mem i li  (* test if (i) is prime *)
              then (fold_step remove li (i*i) n i)
              else li)
	    li 2 (pred limit)
;;


let rec sum = function
  | [] -> 0
  | hd::tl -> hd + sum tl
;;
  


time (fun () -> (sum (primes 2000))) 1
