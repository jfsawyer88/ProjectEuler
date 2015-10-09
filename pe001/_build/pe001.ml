(** Project euler **)
(** Multiples of 3 and 5 **)

(** timing function **)
(*#load "unix.cma";;*)
(*let time f =
  let t = Unix.gettimeofday () in
  let res = f () in
  Printf.printf "Execution time: %f seconds\n"
                (Unix.gettimeofday () -. t);
  Printf.printf "Output: %d" res;
;;
 *)


let time f number =
  let t = Unix.gettimeofday () in
  let rec loop i =
    if i = number then f ()
    else
      begin
	ignore (f ());
	loop (i + 1)
      end
  in
  let res = loop 0 in
  Printf.printf "Execution time: %f seconds\n"
                (Unix.gettimeofday () -. t);
  Printf.printf "Output: %d" res;
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

(*
Printf.printf "Input: ";;
let rec main () =
  let input = input_line stdin in
  if input = "quit" then Printf.printf "Good Bye."
  else
    try
      time (fun ()-> (pe001 (int_of_string input)));
    with 
      Failure "int_of_string" -> Printf.printf "Error: Invalid input"
;;

main ()
*)

time (fun () -> (pe001 1000)) 1000000;;
print_newline ()
