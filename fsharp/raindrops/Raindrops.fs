module Raindrops

let convert (number: int): string =
    [
        if number % 3 = 0 then Some "Pling"
        if number % 5 = 0 then Some "Plang"
        if number % 7 = 0 then Some "Plong"
    ]
    |> List.choose id
    |> function
        | [] -> string number
        | xs -> String.concat "" xs
