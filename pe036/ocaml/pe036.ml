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

let reverse_digits n =
  let rec aux m rev =
    if m = 0 then rev
    else aux (m / 10) ((rev * 10) + (m mod 10))
  in
  aux n 0
;;

let make_room n =
  let rec aux m res =
    if m = 0 then res
    else aux (m / 10) (res * 10)
  in
  aux n 1
;;

let is_palindrome n b =
  let rec aux m rev =
    if m = 0 then rev
    else aux (m / b) ((rev * b) +  (m mod b))
  in
  n = aux n 0
;;

let make_palindrome n odd =
  if odd then (make_room n) * (n / 10) + reverse_digits n
  else (make_room n) * n + reverse_digits n
;;

let pe036 lim =
  let rec aux n res odd =
    let pal = make_palindrome n odd in
    if pal > lim then res
    else
      if is_palindrome pal 2 then
	aux (n + 1) (res + pal) odd
      else
	aux (n + 1) res odd
  in
  aux 0 (aux 0 0 true) false
;;

time (fun () -> pe036 1000000) 1
