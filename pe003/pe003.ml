(** Project Euler **)
(** Largest Prime Factor **)


(** timing function **)
#load "unix.cma";;
let time f =
  let t = Unix.gettimeofday () in
  let res = f () in
  Printf.printf "Execution time: %f seconds\n"
                (Unix.gettimeofday () -. t);
  Printf.printf "Output: %d\n" res;
;;


(** divide out all factors of n from a **)
let rec ddiv a n =
  if 0 <> a mod n then a
  else
    ddiv (a / n) n
;;


let pe003 n =
  let rec loop i n =
    if i = n then n
    else
      loop (i + 1) (ddiv n i)
  in
  loop 2 n
;;


time (fun () -> (pe003 600851475143))

