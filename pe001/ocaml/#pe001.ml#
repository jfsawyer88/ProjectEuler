(** Project euler **)
(** Multiples of 3 and 5 **)


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



let pe001 n =
  let rec loop i res =
    if i = n then res
    else
      if (0 = i mod 3) || (0 = i mod 5) then
	loop (i + 1) (res + i)
      else
	loop (i + 1) res
  in
  loop 1 0;
;;


time (fun () -> (pe001 1000)) 1;;
