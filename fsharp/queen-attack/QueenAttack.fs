module QueenAttack

let create (position: int * int) =
    let row, col = position
    row >= 0 && row <= 7 && col >= 0 && col <= 7

let canAttack (black: int * int) (white: int * int) =
    let a, b = black
    let c, d = white
    (a = c || b = d) || abs(a - c) = abs(b - d)
