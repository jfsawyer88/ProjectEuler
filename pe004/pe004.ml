(** Project Euler **)
(** Largest Palindrome Product **)


(** timing function **)
(*#load "unix.cma";;*)
let time f number =
  let t = Unix.gettimeofday () in  (** get the current time **)
  let time_unit = match number with
    | 1 -> ""
    | 1000 -> "milli-"
    | 1000000 -> "micro-"
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


(** reverses digits of input **)
let reverse_int n =
  let rec loop n rev =
    if !n = 0 then
      !rev
    else
      begin
	rev := (!rev * 10) + (!n mod 10);
	n := !n / 10;
	loop n rev
      end
  in
  loop (ref n) (ref 0)
;;


(** determines if input is a palindrome **)
let is_palindrome n =
  n = reverse_int n
;;


let pe004 () =
  let rec loop a b largest =
    if a < 100 then largest              (** terminate when a < 100 **)
    else
      if b < a then                      (** wlog assume a <= b **)
	loop (a - 1) 999 largest
      else
	if (a * b) < largest then        (** move on to next a value **)
	  loop (a - 1) 999 largest       (** product drops below largest **)
	else
	  if is_palindrome (a * b) then
	    loop a (b - 1) (a * b)
	  else
	    loop a (b - 1) largest
  in
  loop 999 999 0
;;

time pe004 1000000
