(** Project Euler **)
(** Double-base Palindromes **)


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


let is_palindrome n b =
  let rec aux m rev =
    if m = 0 then rev
    else aux (m / b) ((rev * b) +  (m mod b))
  in
  n = aux n 0
;;

let is_double_palindrome n =
  (is_palindrome n 2) && (is_palindrome n 10)
;;

let pe036 limit =
  let rec aux i res =
    if i >= limit then res
    else
      if is_double_palindrome i then aux (i + 2) (res + i)
      else aux (i + 2) res
  in
  aux 1 0
;;

time (fun () -> pe036 1000000) 1
