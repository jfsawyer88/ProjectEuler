(** Project Euler **)
(** Maximum Path Sum I **)

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


let scan_max l =
  let rec aux a l out =
    match l with
    |    []    -> out
    | hd :: tl -> aux hd tl ((max a hd) :: out)
  in
  aux (List.hd l) (List.tl l) []
;;


let expand v n =
  let rec aux n out =
    if n = 0 then out
    else aux (n - 1) (v :: out)
  in
  aux n []
;;


let foo a b =
  List.rev_map2 (+) (scan_max a) b
;;


let process_tri f =
  let sep = Str.regexp " " in
  let chan = open_in f in
  let rec aux out =
    try
      aux ((List.rev_map int_of_string (Str.split sep (input_line chan))) :: out)
    with
    | End_of_file -> (close_in chan; out)
  in
  aux []
;;

let tri = process_tri "tri.txt";;

let pe067 () =
  List.hd (List.fold_left foo (expand 0 (1 + List.length tri)) tri)
;;

time pe067 1000;;
