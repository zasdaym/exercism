module TisburyTreasureHunt

open System

let getCoordinate (line: string * string) : string =
    match line with
    | _, coordinate -> coordinate

let convertCoordinate (coordinate: string) : int * char =
    let x =
        match Int32.TryParse(coordinate[0].ToString()) with
        | true, parsed -> parsed
        | false, _ -> 0

    (x, coordinate[1])

let compareRecords (azarasData: string * string) (ruisData: string * (int * char) * string) : bool =
    let _, azarasCoordinate = azarasData

    let _, ruisCoordinate, _ = ruisData

    convertCoordinate azarasCoordinate = ruisCoordinate

let createRecord
    (azarasData: string * string)
    (ruisData: string * (int * char) * string)
    : (string * string * string * string) =
    let treasure, azarasCoordinate = azarasData
    let location, ruisCoordinate, quadrant = ruisData

    if convertCoordinate azarasCoordinate = ruisCoordinate then
        azarasCoordinate, location, quadrant, treasure
    else
        "", "", "", ""
