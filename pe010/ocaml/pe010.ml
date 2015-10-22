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


(** returns the integer square root of n **)
(** i.e. floor(sqrt(n)) **)
let  isqrt n =
  let rec loop xk = 
    let xk1 = (xk + (n / xk)) / 2 in  (** computs the next x value **)
    if 1 >= (xk - xk1) then xk1       (** compares it with our current x value **)
    else loop xk1
  in
  loop n
;;


let rec  range a b i =
  if a = b then []
  else a :: range (a + i) b i
;;


let pe010 n =
  let r = isqrt n in
  let v = List.map (( / ) n) (range 1 (r + 1) 1) in
  let v = List.append v (range (r - 1) 0 (-1)) in
  let s = Hashtbl.create (List.length v) in
  let sfind = Hashtbl.find s in
  let sadd  = Hashtbl.add  s in
  let srepl = Hashtbl.replace s in
  for i = 0 to (List.length v) - 1 do
    sadd
      (List.nth v i)
		(
		  (
		    ((List.nth v i) * ((List.nth v i) + 1)) / 2
		  ) - 1
		)
  done;
  for p = 2 to r do
    let sp = sfind (p - 1) in
    if (sfind p) > sp then
      let p2 = p * p in
      let rec loop i =
	let vi = List.nth v i in
	if vi < p2 then ()
	else
	  begin
	    srepl
	      vi
	      ((sfind vi) - (p * ((sfind (vi / p)) - sp))
	      );
	    loop (succ i)
	  end
      in
      loop 0
  done;
  Hashtbl.find s n
;;


time (fun () -> (pe010 2000000)) 1
