(** Project Euler **)
(** Even Fibonacci Numbers **)


(** timing function **)
(*#load "unix.cma";;*)
let time f =
  let t = Unix.gettimeofday () in
  let res = f () in
  Printf.printf "Execution time: %f seconds\n"
                (Unix.gettimeofday () -. t);
  Printf.printf "Output: %d" res;
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


time (fun () -> (pe002 1000));;
print_newline ()
