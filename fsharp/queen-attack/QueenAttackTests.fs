module QueenAttackTests

open FsUnit.Xunit
open Xunit

open QueenAttack

[<Fact>]
let ``Queen with a valid position`` () =
    create (2, 2) |> should equal true

[<Fact>]
let ``Queen must have positive row`` () =
    create (-2, 2) |> should equal false

[<Fact>]
let ``Queen must have row on board`` () =
    create (8, 4) |> should equal false

[<Fact>]
let ``Queen must have positive column`` () =
    create (2, -2) |> should equal false

[<Fact>]
let ``Queen must have column on board`` () =
    create (4, 8) |> should equal false

[<Fact>]
let ``Cannot attack`` () =
    let whiteQueen = (2, 4)
    let blackQueen = (6, 6)
    canAttack blackQueen whiteQueen |> should equal false

[<Fact>]
let ``Can attack on same row`` () =
    let whiteQueen = (2, 4)
    let blackQueen = (2, 6)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Can attack on same column`` () =
    let whiteQueen = (4, 5)
    let blackQueen = (2, 5)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Can attack on first diagonal`` () =
    let whiteQueen = (2, 2)
    let blackQueen = (0, 4)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Can attack on second diagonal`` () =
    let whiteQueen = (2, 2)
    let blackQueen = (3, 1)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Can attack on third diagonal`` () =
    let whiteQueen = (2, 2)
    let blackQueen = (1, 1)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Can attack on fourth diagonal`` () =
    let whiteQueen = (1, 7)
    let blackQueen = (0, 6)
    canAttack blackQueen whiteQueen |> should equal true

[<Fact>]
let ``Cannot attack if falling diagonals are only the same when reflected across the longest falling diagonal`` () =
    let whiteQueen = (4, 1)
    let blackQueen = (2, 5)
    canAttack blackQueen whiteQueen |> should equal false
