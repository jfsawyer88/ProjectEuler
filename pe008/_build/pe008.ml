(** Project Euler **)
(** Largest Product in a Series **)


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
  let res = loop 0 in              (** keep the result to print **)
  Printf.printf "Execution time: %f %sseconds\n"
                (Unix.gettimeofday () -. t)
		time_unit;
  Printf.printf "Output: %d\n" res;
;;

let series_str = "
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450
"
;;

let process_series series =
  let rec loop i res =
    if i = String.length series then res
    else
      match series.[i] with
      | '0' -> loop (i + 1) (Array.append res [|0|])
      | '1' -> loop (i + 1) (Array.append res [|1|])
      | '2' -> loop (i + 1) (Array.append res [|2|])
      | '3' -> loop (i + 1) (Array.append res [|3|])
      | '4' -> loop (i + 1) (Array.append res [|4|])
      | '5' -> loop (i + 1) (Array.append res [|5|])
      | '6' -> loop (i + 1) (Array.append res [|6|])
      | '7' -> loop (i + 1) (Array.append res [|7|])
      | '8' -> loop (i + 1) (Array.append res [|8|])
      | '9' -> loop (i + 1) (Array.append res [|9|])
      |  _  -> loop (i + 1)  res
  in
  loop 0 [||]
;;

let series_int = process_series series_str;;


let pe008 series =
  let rec loop i res=
    if (i + 13) =  (Array.length series) then res
    else
      loop (i + 1)
	   (max res
		(Array.fold_left 
		   ( * ) 
		   1
		   (Array.sub series_int i 13)
		)
	   )
  in
  loop 0 0
;;


time (fun () -> (pe008 series_int)) 1000000
