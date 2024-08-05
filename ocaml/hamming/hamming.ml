type nucleotide = A | C | G | T

let combine a b c = a + if b = c then 0 else 1

let hamming_distance a b =
  match (List.length a, List.length b) with
  | 0, 0 -> Ok 0
  | 0, _ -> Error "left strand must not be empty"
  | _, 0 -> Error "right strand must not be empty"
  | x, y when x != y -> Error "left and right strands must be of equal length"
  | _ -> Ok (List.fold_left2 combine 0 a b)
